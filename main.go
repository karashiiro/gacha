package main

import (
	"errors"
	"math/rand"
	"time"
)

var rng *rand.Rand

func checkRoll(drops []Drop, val float32) (*Drop, error) {
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
	_, err = NewDatabase(sugar)
	if err != nil {
		sugar.Errorw("couldn't connect to database, aborting")
		panic(err)
	}

	// Initialize randomizer with current time
	rngSource := rand.NewSource(time.Now().Unix())
	rng = rand.New(rngSource)

	sugar.Infow("application started")

	sugar.Infow("random number generated",
		"number", rng.Float32(),
	)
}
