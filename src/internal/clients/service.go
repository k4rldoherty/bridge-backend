// Package clients - Service Layer
package clients

import (
	"context"
	"database/sql"
	"encoding/json"
	"log/slog"

	"github.com/k4rldoherty/brige-backend/src/internal/db"
)

type Service interface {
	GetClients(ctx context.Context) ([]db.Client, error)
	AddClient(ctx context.Context, c []byte) (db.Client, error)
}

type svc struct {
	repo db.Querier
}

type clientDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	JoinCode string `json:"join_code"`
	LogoURL  string `json:"logo_url"`
}

func NewService(q db.Querier) Service {
	return &svc{
		repo: q,
	}
}

// the service is responsible for handling any business logic related to the request
func (s *svc) GetClients(ctx context.Context) ([]db.Client, error) {
	c, err := s.repo.GetClients(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *svc) AddClient(ctx context.Context, data []byte) (db.Client, error) {
	// Create a client object by unmarshalling the json
	c := clientDTO{}
	if err := json.Unmarshal(data, &c); err != nil {
		slog.Error("failed to create client object", "error", err, "location", "service.AddClient")
		return db.Client{}, err
	}

	// Parse the db query params from the data
	params := db.AddClientParams{
		Name:     c.Name,
		Email:    c.Email,
		JoinCode: c.JoinCode,
		LogoUrl:  sql.NullString{String: c.LogoURL, Valid: true},
	}

	// Call the repo to add the client to the database
	addedClient, err := s.repo.AddClient(ctx, params)
	if err != nil {
		return db.Client{}, err
	}
	return addedClient, nil
}
