// Package clients - Service Layer
package clients

import (
	"context"

	"github.com/k4rldoherty/brige-backend/src/internal/db"
)

type Service interface {
	GetClients(ctx context.Context) ([]db.Client, error)
}

type svc struct {
	repo db.Querier
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
