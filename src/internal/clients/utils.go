package clients

import (
	"net/http"

	"github.com/k4rldoherty/brige-backend/src/internal/db"
	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	"github.com/k4rldoherty/brige-backend/src/internal/utils"
)

type CreateClientDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	JoinCode string `json:"join_code"`
	LogoURL  string `json:"logo_url"`
}

type UpdateClientDTO struct {
	ID      int32  `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}

func NewService(q db.Querier, l *logger.Logger) Service {
	return &svc{
		repo:   q,
		logger: l,
	}
}

func (c CreateClientDTO) ValidateInput() *utils.APIError {
	if c.Name == "" {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "name is required",
		}
	}
	if c.Email == "" {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "email is required",
		}
	}
	if c.JoinCode == "" {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "join_code is required",
		}
	}
	return nil
}

func (c UpdateClientDTO) ValidateInput() *utils.APIError {
	if c.ID < 1 {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "id is required and must be a valid number greater than 0, and inside the int32 range",
		}
	}
	if c.Email == "" {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "email is required",
		}
	}
	if c.Name == "" {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "name is required",
		}
	}
	return nil
}
