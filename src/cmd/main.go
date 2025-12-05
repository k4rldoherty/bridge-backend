package main

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Set slog as the default logger
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(l)

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("failed to load .env", "error", err)
		os.Exit(1)
	}

	// Get environment variables
	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")

	dbQueries, err := sql.Open("postgres", dbURL)
	if err != nil {
		slog.Error("failed to open db", "error", err, "location", "main.main")
		os.Exit(1)
	}

	defer func() {
		if err := dbQueries.Close(); err != nil {
			slog.Error("failed to close db", "error", err, "location", "main.main")
			os.Exit(1)
		}
	}()

	slog.Info("db connected successfully")

	cfg := config{
		addr: ":" + port,
		db: dbConfig{
			dbQueries: dbQueries,
		},
	}

	api := &app{cfg: cfg}

	if err := api.start(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
