package main

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/karashiiro/gacha/ent"
	"github.com/karashiiro/gacha/ent/drop"
	"github.com/karashiiro/gacha/ent/series"
	"github.com/karashiiro/gacha/message"
	"go.uber.org/zap"
)

// Database represents an ent database with a Redis cache-aside.
type Database struct {
	edb   *ent.Client
	rdb   *cache.Cache
	sugar *zap.SugaredLogger
}

// NewDatabase creates a new database connection.
func NewDatabase(sugar *zap.SugaredLogger) (*Database, error) {
	// Connect to database
	var db *sql.DB
	var err error
	for db == nil {
		db, err = sql.Open("mysql", os.Getenv("GACHA_MYSQL_CONNECTION_STRING"))
		if err != nil {
			sugar.Errorw("failed to connect to database, retrying in 5 seconds",
				"error", err,
			)
			time.Sleep(5 * time.Second)
		}
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	driver := entsql.OpenDB("mysql", db)

	var dbCreateQuery = "CREATE DATABASE IF NOT EXISTS gacha CHARACTER SET = 'utf8';"
	_, err = db.Exec(dbCreateQuery)
	for err != nil {
		sugar.Errorw("failed to create database, retrying in 5 seconds",
			"error", err,
		)
		time.Sleep(5 * time.Second)
		_, err = db.Exec(dbCreateQuery)
	}

	_, err = db.Exec("USE gacha;")
	if err != nil {
		sugar.Errorw("failed to switch to database",
			"error", err,
		)
		return nil, err
	}

	edb := ent.NewClient(ent.Driver(driver))

	// Run auto-migration
	if err := edb.Schema.Create(context.Background()); err != nil {
		sugar.Errorw("failed to auto-migrate schema",
			"error", err,
		)
		return nil, err
	}

	// Connect to Redis server
	r := redis.NewClient(&redis.Options{
		Addr: os.Getenv("GACHA_REDIS_LOCATION"),
	})

	c := cache.New(&cache.Options{
		Redis: r,
	})

	d := &Database{
		edb:   edb,
		rdb:   c,
		sugar: sugar,
	}

	return d, nil
}

// GetDropTable returns the drop table associated with the specified series name.
func (d *Database) GetDropTable(name string) ([]ent.Drop, error) {
	var dropTable []ent.Drop

	err := d.rdb.Get(context.Background(), name, &dropTable)
	if err != nil {
		d.sugar.Warnw("cache miss occurred, fetching from database",
			"name", name,
		)

		ctx := context.Background()

		rows, err := d.edb.Drop.
			Query().
			Where(drop.HasSeriesWith(series.NameEQ(name))).
			All(ctx)
		if err != nil {
			return nil, err
		}

		// Copy the DB rows into a slice
		drops := make([]ent.Drop, len(rows))
		for i, row := range rows {
			drops[i] = *row
		}

		// Cache the slice
		err = d.rdb.Set(&cache.Item{
			Ctx:   ctx,
			Key:   name,
			Value: drops,
			TTL:   24 * time.Hour,
		})
		if err != nil {
			return nil, err
		}

		dropTable = drops
	}

	return dropTable, nil
}

// Adds a new series to the database, or returns an existing one.
func (d *Database) addSeries(ctx context.Context, seriesName string) (*ent.Series, error) {
	sr, err := d.edb.Series.Query().
		Where(series.NameEQ(seriesName)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	if len(sr) != 0 {
		return sr[0], nil
	}

	return d.edb.Series.Create().
		SetName(seriesName).
		Save(ctx)
}

// DeleteDropTable deletes a series and all of its children.
func (d *Database) DeleteDropTable(seriesName string) error {
	ctx := context.Background()

	_, err := d.edb.Drop.Delete().
		Where(drop.HasSeriesWith(series.NameEQ(seriesName))).
		Exec(ctx)
	if err != nil {
		return err
	}

	_, err = d.edb.Series.Delete().
		Where(series.NameEQ(seriesName)).
		Exec(ctx)
	if err != nil {
		return err
	}

	err = d.rdb.Delete(ctx, seriesName)

	return err
}

// AddDropTable adds a single drop table to the database. The rates of each inserted drop must sum to 1.0.
func (d *Database) SetDropTable(seriesName string, drops []message.DropInsert) error {
	ctx := context.Background()

	sr, err := d.addSeries(ctx, seriesName)
	if err != nil {
		return err
	}

	agg := float32(0)
	builders := make([]*ent.DropCreate, len(drops))
	for i, di := range drops {
		agg += di.Rate
		builders[i] = d.edb.Drop.Create().
			SetObjectID(di.ObjectID).
			SetRate(di.Rate).
			SetSeriesID(sr.ID)
	}

	if agg != 1.0 {
		return errors.New("series rates must sum to 1.0")
	}

	_, err = d.edb.Drop.Delete().
		Where(drop.HasSeriesWith(series.NameEQ(seriesName))).
		Exec(ctx)
	if err != nil {
		return err
	}

	dbDrops, err := d.edb.Drop.CreateBulk(builders...).
		Save(ctx)
	if err != nil {
		return err
	}

	// Copy the DB rows into a slice
	dbDropValues := make([]ent.Drop, len(dbDrops))
	for i, dr := range dbDrops {
		dbDropValues[i] = *dr
	}

	d.rdb.Set(&cache.Item{
		Ctx:   ctx,
		Key:   seriesName,
		Value: dbDropValues,
		TTL:   24 * time.Hour,
	})

	return err
}
