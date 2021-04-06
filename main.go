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
	rngSource := rand.NewSource(time.Now().Unix())
	rng = rand.New(rngSource)

	sugar.Infow("application started")

	sugar.Infow("random number generated",
		"number", rng.Float32(),
	)

	rows, err := db.GetDropTable("test")
	if err != nil {
		sugar.Errorw("failed to get rows",
			"error", err,
		)
	}

	roll, err := checkRoll(rows, rng.Float32())
	if err != nil {
		sugar.Errorw("gacha roll failed",
			"error", err,
		)
	}

	sugar.Infof("Rolled %v", roll)
}
