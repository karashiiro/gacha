package main

import (
	"math/rand"
	"time"
)

var rng *rand.Rand

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
