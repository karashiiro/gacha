package main

import (
	"errors"
	"math/rand"
	"time"

	"github.com/karashiiro/gacha/ent"
)

var rng *rand.Rand

func checkRoll(drops []ent.Drop, val float32) (*ent.Drop, error) {
	agg := float32(0)
	for _, drop := range drops {
		agg += drop.Rate
		if agg > 1.0 {
			break
		}
		if agg > val {
			return &drop, nil
		}
	}
	return nil, errors.New("drop rates do not sum to 1.0")
}

func main() {
	// Set up logging
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// Connect to database
	db, err := NewDatabase(sugar)
	if err != nil {
		sugar.Errorw("couldn't connect to database, aborting")
		panic(err)
	}
	defer db.edb.Close()

	// Initialize randomizer with current time
	rngSource := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(rngSource)

	sugar.Infow("application started")

	rows, err := db.GetDropTable("test")
	if err != nil {
		sugar.Errorw("failed to get rows",
			"error", err,
		)
	}

	testValue := rng.Float32()
	sugar.Infow("random number generated",
		"number", testValue,
	)

	roll, err := checkRoll(rows, testValue)
	if err != nil {
		sugar.Errorw("gacha roll failed",
			"error", err,
		)
	}

	sugar.Infof("Rolled %v", roll)
}
