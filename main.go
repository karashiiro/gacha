package main

import (
	"os"

	"go.uber.org/zap"
)

// NewLogger creates a file/console logger.
func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = append(cfg.OutputPaths, "log/gacha.log")

	// Create log directory if it does not exist
	_, err := os.Stat("log/")
	if os.IsNotExist(err) {
		os.Mkdir("log/", 0644)
	}

	return cfg.Build()
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
	_, err = NewDatabase()
	if err != nil {
		sugar.Errorw("couldn't connect to database, aborting")
		panic(err)
	}

	sugar.Infow("application started")
}
