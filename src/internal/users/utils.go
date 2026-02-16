package users

import (
	"net/http"

	"github.com/k4rldoherty/brige-backend/src/internal/db"
	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	"github.com/k4rldoherty/brige-backend/src/internal/utils"
)

type CreateUserDTO struct {
	ClientID int32  `json:"client_id"`
	RoleID   int32  `json:"role_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewService(q db.Querier, l *logger.Logger) Service {
	return &svc{
		repo:   q,
		logger: l,
	}
}

func (c CreateUserDTO) ValidateInput() *utils.APIError {
	if c.ClientID < 1 {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "role_id is required and must be a valid number greater than 0, and inside the int32 range",
		}
	}
	if c.RoleID < 1 {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "role_id is required and must be a valid number greater than 0, and inside the int32 range",
		}
	}
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
	if c.Password == "" {
		return &utils.APIError{
			Status:  http.StatusBadRequest,
			Message: "password is required",
		}
	}
	return nil
}
