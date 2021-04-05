package main

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
