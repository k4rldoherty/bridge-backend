package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/k4rldoherty/brige-backend/src/internal/clients"
	"github.com/k4rldoherty/brige-backend/src/internal/db"
)

type app struct {
	cfg config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dbQueries *sql.DB
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
			log.Println(err)
		}
	})

	// Create a handler, all handlers must be of the signature w http.ResponseWriter, r *http.Request
	clientsHandler := clients.NewHandler(clients.NewService(db.New(a.cfg.db.dbQueries)))
	r.Get("/clients", clientsHandler.GetClients)
	r.Post("/clients", clientsHandler.AddClient)
	r.Put("/clients", clientsHandler.UpdateClient)

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

	log.Printf("Starting server on port %s", a.cfg.addr)

	return srv.ListenAndServe()
}
