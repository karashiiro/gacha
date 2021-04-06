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
	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	driver := entsql.OpenDB("mysql", db)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS gacha CHARACTER SET = 'utf8';")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("USE gacha;")
	if err != nil {
		return nil, err
	}

	edb := ent.NewClient(ent.Driver(driver))

	// Run auto-migration
	if err := edb.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	// Connect to Redis server
	r := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_LOCATION"),
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
			TTL:   30 * 24 * time.Hour,
		})
		if err != nil {
			return nil, err
		}

		dropTable = drops
	}

	return dropTable, nil
}

// AddSeries constructs a new series and inserts it into the database.
func (d *Database) AddSeries(name string) error {
	_, err := d.edb.Series.Create().
		SetName(name).
		Save(context.Background())
	return err
}

// DeleteSeries deletes a series and all of its children.
func (d *Database) DeleteSeries(name string) error {
	ctx := context.Background()

	_, err := d.edb.Drop.Delete().
		Where(drop.HasSeriesWith(series.NameEQ(name))).
		Exec(ctx)
	if err != nil {
		return err
	}

	_, err = d.edb.Series.Delete().
		Where(series.NameEQ(name)).
		Exec(ctx)

	return err
}

// DropInsert is an insert request for a new drop.
type DropInsert struct {
	ObjectID uint32  `json:"object_id"`
	Rate     float32 `json:"rate"`
}

// AddDropTable adds a single drop table to the database. The rates of each inserted drop must sum to 1.0.
func (d *Database) SetDropTable(seriesName string, drops []DropInsert) error {
	ctx := context.Background()

	sr, err := d.edb.Series.Query().
		Where(series.NameEQ(seriesName)).
		First(ctx)
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

	_, err = d.edb.Drop.CreateBulk(builders...).
		Save(ctx)

	return err
}
