package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/k4rldoherty/brige-backend/src/internal/clients"
	"github.com/k4rldoherty/brige-backend/src/internal/db"
	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	"github.com/k4rldoherty/brige-backend/src/internal/utils"
)

type app struct {
	cfg    config
	logger *logger.Logger
	db     *sql.DB
}

type config struct {
	addr    string
	connStr string
}

// sets up the chi router and returns a handler
func (a *app) mount() http.Handler {
	r := chi.NewRouter()

	// a base middleware stack
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	// basic health check
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("All good!\n")); err != nil {
			a.logger.Error("failed to write response", "error", err, "location", "cmd/api.go")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// TODO: add auth
	r.Delete("/admin/reset", func(w http.ResponseWriter, r *http.Request) {
		res, err := a.db.Exec("DELETE FROM clients")
		if err != nil {
			a.logger.Error("failed to write response", "error", err, "location", "cmd/api.go")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if n, err := res.RowsAffected(); err != nil {
			a.logger.Error("failed to get rows affected", "error", err, "location", "cmd/api.go")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if n == 0 {
			a.logger.Error("no rows deleted", "error", "no rows deleted", "location", "cmd/api.go")
			http.Error(w, "no rows deleted", http.StatusInternalServerError)
			return
		}
		utils.Write(w, http.StatusOK, "Database reset")
	})

	// Create a handler, all handlers must be of the signature w http.ResponseWriter, r *http.Request
	clientsHandler := clients.NewHandler(clients.NewService(db.New(a.db), a.logger), a.logger)
	r.Get("/clients", clientsHandler.GetClients)
	r.Post("/clients", clientsHandler.AddClient)
	r.Put("/clients", clientsHandler.UpdateClient)
	r.Delete("/clients/{id}", clientsHandler.DeleteClient)

	return r
}

// created a new server struct and starts up a server on the port specified
func (a *app) start(h http.Handler) error {
	srv := &http.Server{
		Addr:         a.cfg.addr,
		Handler:      h,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}

	a.logger.Info("server started", "location", "cmd/api.go")

	return srv.ListenAndServe()
}
