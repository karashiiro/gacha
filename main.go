package main

import (
	"os"

	"go.uber.org/zap"
)

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
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Application started.")
}
