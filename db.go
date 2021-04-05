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
	gz.sugar.Infow(msg, keysAndValues)
}

func (gz *GormZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	gz.sugar.Warnw(msg, keysAndValues)
}

func (gz *GormZapLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	gz.sugar.Errorw(msg, keysAndValues)
}

func (gz *GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		gz.sugar.Errorf("%s\n[%.3fms] [rows:%v] %s", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}

func NewDatabase(sugar *zap.SugaredLogger) (*Database, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_LOCATION")), &gorm.Config{
		Logger: &GormZapLogger{sugar: sugar},
	})
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
