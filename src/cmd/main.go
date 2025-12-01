package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")

	cfg := config{
		addr: ":" + port,
		db: dbConfig{
			connString: dbURL,
		},
	}

	api := &app{cfg: cfg}

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(l)

	if err := api.start(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
