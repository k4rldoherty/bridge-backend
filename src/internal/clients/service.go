// Package clients - Service Layer
package clients

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/k4rldoherty/brige-backend/src/internal/db"
	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	"github.com/k4rldoherty/brige-backend/src/internal/utils"
)

func NewService(q db.Querier, l *logger.Logger) Service {
	return &svc{
		repo:   q,
		logger: l,
	}
}

func (s *svc) GetClients(ctx context.Context) ([]db.Client, *utils.APIError) {
	c, err := s.repo.GetClients(ctx)
	if err != nil {
		return nil, &utils.APIError{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	return c, nil
}

func (s *svc) AddClient(ctx context.Context, data []byte) (db.Client, *utils.APIError) {
	c := CreateClientDTO{}
	if err := json.Unmarshal(data, &c); err != nil {
		s.logger.Error("failed to create client object request body", "error", err, "location", "service.AddClient")
		return db.Client{}, &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if err := c.ValidateInput(); err != nil {
		s.logger.Error("failed to validate client input", "error", err.Message, "location", "service.AddClient")
		return db.Client{}, err
	}

	params := db.AddClientParams{
		Name:     c.Name,
		Email:    c.Email,
		JoinCode: c.JoinCode,
		LogoUrl:  utils.ToNullString(c.LogoURL),
	}

	addedClient, err := s.repo.AddClient(ctx, params)
	if err != nil {
		return db.Client{}, &utils.APIError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return addedClient, nil
}

func (s *svc) UpdateClient(ctx context.Context, data []byte) (db.Client, *utils.APIError) {
	c := UpdateClientDTO{}
	if err := json.Unmarshal(data, &c); err != nil {
		s.logger.Error("failed to create client object from request body", "error", err, "location", "service.UpdateClient")
		return db.Client{}, &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if err := c.ValidateInput(); err != nil {
		s.logger.Error("failed to validate client input", "error", err, "location", "service.UpdateClient")
		return db.Client{}, err
	}
	params := db.UpdateClientParams{
		ID:      c.ID,
		Name:    c.Name,
		Email:   c.Email,
		LogoUrl: utils.ToNullString(c.LogoURL),
	}

	updatedClient, err := s.repo.UpdateClient(ctx, params)
	if err != nil {
		return db.Client{}, &utils.APIError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return updatedClient, nil
}

func (s *svc) DeleteClient(ctx context.Context, data string) *utils.APIError {
	id, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		s.logger.Error("failed to parse client id from request body", "error", err, "location", "service.DeleteClient")
		return &utils.APIError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	_, err = s.repo.DeleteClient(ctx, int32(id))
	if err != nil {
		s.logger.Error("failed to delete client", "error", err, "location", "service.DeleteClient")
		return &utils.APIError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}
