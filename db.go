package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database represents a GORM database with a Redis cache-aside.
type Database struct {
	db    *gorm.DB
	rdb   *redis.Client
	sugar *zap.SugaredLogger
}

// GormZapLogger is a GORM adapter for zap.
type GormZapLogger struct {
	sugar *zap.SugaredLogger
}

func (gz *GormZapLogger) LogMode(level logger.LogLevel) logger.Interface {
	return gz
}

func (gz *GormZapLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	gz.sugar.Infof(msg, keysAndValues)
}

func (gz *GormZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	gz.sugar.Warnf(msg, keysAndValues)
}

func (gz *GormZapLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	gz.sugar.Errorf(msg, keysAndValues)
}

func (gz *GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		gz.sugar.Errorf("%s\n[%.3fms] [rows:%v] %s", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}

func NewDatabase(sugar *zap.SugaredLogger) (*Database, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_CONNECTION_STRING")), &gorm.Config{
		Logger: &GormZapLogger{sugar: sugar},
	})
	if err != nil {
		return nil, err
	}

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

func (d *Database) GetDropTable(name string) ([]Drop, error) {
	var dropTable []Drop

	val, err := d.rdb.Get(context.Background(), name).Result()
	if err == redis.Nil {
		table := d.db.Table(name)
		table.AutoMigrate(&Drop{})

		rows, err := table.Rows()
		if err != nil {
			d.sugar.Errorw("failed to fetch table rows",
				"name", name,
			)

			return nil, err
		}

		// Copy the DB rows into a slice
		var rowCount int64
		table.Count(&rowCount)
		drops := make([]Drop, rowCount)
		i := 0
		for rows.Next() {
			err = rows.Scan(&drops[i])
			if err != nil {
				d.sugar.Error("failed to copy rows to slice")
				return nil, err
			}
		}

		// Cache the slice
		err = d.rdb.Set(context.Background(), name, drops, 0).Err()
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
