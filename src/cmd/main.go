package main

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	_ "github.com/lib/pq"
)

func main() {
	logger := logger.NewLogger()

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("failed to load .env", "error", err, "location", "main.main")
		os.Exit(1)
	}

	// Get environment variables
	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")

	if dbURL == "" || port == "" {
		logger.Error("missing environment variables", "location", "main.main")
		os.Exit(1)
	}

	dbQueries, err := sql.Open("postgres", dbURL)
	if err != nil {
		logger.Error("failed to open db", "error", err, "location", "main.main")
		os.Exit(1)
	}

	defer func() {
		if err := dbQueries.Close(); err != nil {
			logger.Error("failed to close db", "error", err, "location", "main.main")
			os.Exit(1)
		}
	}()

	logger.Info("db connected successfully", "location", "main.main")

	cfg := config{
		addr:    ":" + port,
		connStr: dbURL,
	}

	api := &app{
		cfg:    cfg,
		logger: logger,
		db:     dbQueries,
	}

	if err := api.start(api.mount()); err != nil {
		logger.Error("server failed to start", "error", err, "location", "main.main")
		os.Exit(1)
	}
}
