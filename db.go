package main

import (
	"context"
	"database/sql"
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

// Database represents a ent database with a Redis cache-aside.
type Database struct {
	edb   *ent.Client
	rdb   *cache.Cache
	sugar *zap.SugaredLogger
}

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
			d.sugar.Errorw("failed to fetch table rows",
				"name", name,
			)

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
			d.sugar.Errorw("failed to cache table",
				"name", name,
			)

			return nil, err
		}

		dropTable = drops
	}

	return dropTable, nil
}
