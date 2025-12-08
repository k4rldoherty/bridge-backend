// Package clients - Service Layer
package clients

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/k4rldoherty/brige-backend/src/internal/db"
	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	"github.com/k4rldoherty/brige-backend/src/internal/utils"
)

type Service interface {
	GetClients(ctx context.Context) ([]db.Client, error)
	AddClient(ctx context.Context, d []byte) (db.Client, error)
	UpdateClient(ctx context.Context, d []byte) (db.Client, error)
	DeleteClient(ctx context.Context, d []byte) error
}

type svc struct {
	repo   db.Querier
	logger *logger.Logger
}

func NewService(q db.Querier, l *logger.Logger) Service {
	return &svc{
		repo:   q,
		logger: l,
	}
}

func (s *svc) GetClients(ctx context.Context) ([]db.Client, error) {
	c, err := s.repo.GetClients(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *svc) AddClient(ctx context.Context, data []byte) (db.Client, error) {
	// Create a client object by unmarshalling the json
	c := CreateClientDTO{}
	if err := json.Unmarshal(data, &c); err != nil {
		s.logger.Error("failed to create client object request body", "error", err, "location", "service.AddClient")
		return db.Client{}, err
	}

	if err := c.ValidateInput(); err != nil {
		s.logger.Error("failed to validate client input", "error", err, "location", "service.AddClient")
		return db.Client{}, err
	}

	// Parse the db query params from the data
	params := db.AddClientParams{
		Name:     c.Name,
		Email:    c.Email,
		JoinCode: c.JoinCode,
		LogoUrl:  utils.ToNullString(c.LogoURL),
	}

	// Call the repo to add the client to the database
	addedClient, err := s.repo.AddClient(ctx, params)
	if err != nil {
		return db.Client{}, err
	}
	return addedClient, nil
}

func (s *svc) UpdateClient(ctx context.Context, data []byte) (db.Client, error) {
	c := UpdateClientDTO{}
	if err := json.Unmarshal(data, &c); err != nil {
		s.logger.Error("failed to create client object from request body", "error", err, "location", "service.UpdateClient")
		return db.Client{}, err
	}

	if err := c.ValidateInput(); err != nil {
		s.logger.Error("failed to validate client input", "error", err, "location", "service.UpdateClient")
		return db.Client{}, err
	}

	params := db.UpdateClientParams{
		ID:      int32(c.ID),
		Name:    c.Name,
		Email:   c.Email,
		LogoUrl: utils.ToNullString(c.LogoURL),
	}

	updatedClient, err := s.repo.UpdateClient(ctx, params)
	if err != nil {
		return db.Client{}, err
	}

	return updatedClient, nil
}

func (s *svc) DeleteClient(ctx context.Context, data []byte) error {
	id, err := strconv.Atoi(string(data))
	if err != nil {
		s.logger.Error("failed to parse client id from request body", "error", err, "location", "service.DeleteClient")
		return err
	}
	err = s.repo.DeleteClient(ctx, int32(id))
	if err != nil {
		s.logger.Error("failed to delete client", "error", err, "location", "service.DeleteClient")
		return err
	}
	return nil
}
