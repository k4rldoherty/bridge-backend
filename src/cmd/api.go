package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	cfg config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	connString string
}

// sets up the api
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

	return r
}

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
