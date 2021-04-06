package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/karashiiro/gacha/ent"
	"github.com/karashiiro/gacha/ent/drop"
	"go.uber.org/zap"
)

// Database represents a GORM database with a Redis cache-aside.
type Database struct {
	edb   *ent.Client
	rdb   *redis.Client
	sugar *zap.SugaredLogger
}

func NewDatabase(sugar *zap.SugaredLogger) (*Database, error) {
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

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_LOCATION"),
	})

	d := &Database{
		edb:   edb,
		rdb:   rdb,
		sugar: sugar,
	}

	return d, nil
}

func (d *Database) GetDropTable(name string) ([]ent.Drop, error) {
	var dropTable []ent.Drop

	val, err := d.rdb.Get(context.Background(), name).Result()
	if err == redis.Nil {
		ctx := context.Background()

		rows, err := d.edb.Drop.Query().Where(drop.SeriesEQ(name)).All(ctx)
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
		err = d.rdb.Set(ctx, name, drops, 0).Err()
		if err != nil {
			d.sugar.Errorw("failed to cache table",
				"name", name,
			)

			return nil, err
		}

		dropTable = drops
	}

	err = json.Unmarshal([]byte(val), &dropTable)
	if err != nil {
		d.sugar.Errorw("failed to unmarshal table from Redis",
			"name", name,
		)
	}

	return dropTable, err
}
