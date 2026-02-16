// Package users - handles all modification of users
package users

import (
	"io"
	"net/http"

	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	"github.com/k4rldoherty/brige-backend/src/internal/utils"
)

type handler struct {
	service Service
	logger  *logger.Logger
}

func NewHandler(service Service, logger *logger.Logger) *handler {
	return &handler{service: service, logger: logger}
}

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	d, e := io.ReadAll(r.Body)
	defer utils.CloseRequestBody(r, h.logger)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	if d == nil {
		http.Error(w, "no data provided", http.StatusBadRequest)
		return
	}
	c, err := h.service.AddUser(r.Context(), d)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}
	utils.Write(w, http.StatusOK, c)
}
