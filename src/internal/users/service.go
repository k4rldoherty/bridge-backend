package users

import (
	"context"
	"net/http"

	"github.com/k4rldoherty/brige-backend/src/internal/db"
	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	"github.com/k4rldoherty/brige-backend/src/internal/utils"
)

type Service interface {
	AddUser(ctx context.Context, data []byte) (db.User, *utils.APIError)
}

type svc struct {
	repo   db.Querier
	logger *logger.Logger
}

func (s *svc) AddUser(ctx context.Context, data []byte) (db.User, *utils.APIError) {
	u := CreateUserDTO{}
	err := utils.UnmarshalJSON(data, &u)
	if err != nil {
		return db.User{}, err
	}

	if err = u.ValidateInput(); err != nil {
		return db.User{}, err
	}

	// check role exists
	// check client exists
	// hash password

	params := db.AddUserParams{
		ClientID: utils.ToNullInt32(u.ClientID),
		RoleID:   utils.ToNullInt32(u.RoleID),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password, // change this to hashed password
	}
	user, e := s.repo.AddUser(ctx, params)
	if e != nil {
		return db.User{}, &utils.APIError{
			Status:  http.StatusInternalServerError,
			Message: e.Error(),
		}
	}

	return user, nil
}
