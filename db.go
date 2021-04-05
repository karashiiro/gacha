package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database represents a GORM database with a Redis cache-aside.
type Database struct {
	db    *gorm.DB
	rdb   *redis.Client
	sugar *zap.SugaredLogger
}

func NewDatabase(sugar *zap.SugaredLogger) (*Database, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_LOCATION")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Drop{})

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_LOCATION"),
	})

	d := &Database{
		db:    db,
		rdb:   rdb,
		sugar: sugar,
	}

	return d, nil
}

func (d *Database) GetTable(name string) ([]Drop, error) {
	var dropTable []Drop

	val, err := d.rdb.Get(context.Background(), name).Result()
	if err == redis.Nil {
		dbVal, result := d.db.Get(name)
		if !result {
			d.sugar.Errorw("failed to fetch table",
				"name", name,
			)

			return nil, err
		}

		err = d.rdb.Set(context.Background(), name, dbVal, 0).Err()
		if err != nil {
			d.sugar.Errorw("failed to cache table",
				"name", name,
			)

			return nil, err
		}

		dropTable = dbVal.([]Drop)
	}

	err = json.Unmarshal([]byte(val), &dropTable)
	if err != nil {
		d.sugar.Errorw("failed to unmarshal table from Redis",
			"name", name,
		)
	}

	return dropTable, err
}
