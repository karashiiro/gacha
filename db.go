package main

import (
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database represents a GORM database with a Redis cache-aside.
type Database struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewDatabase() (*Database, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_LOCATION")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Drop{})

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_LOCATION"),
	})

	d := &Database{
		db:  db,
		rdb: rdb,
	}

	return d, nil
}
