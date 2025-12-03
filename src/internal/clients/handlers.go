// Package clients - Handlers
package clients

import (
	"log/slog"
	"net/http"

	"github.com/k4rldoherty/brige-backend/src/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

// the handler is responsible for handling the request and returning a response
func (h *handler) GetClients(w http.ResponseWriter, r *http.Request) {
	c, err := h.service.GetClients(r.Context())
	if err != nil {
		slog.Error("failed to get clients", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(c) == 0 || c == nil {
		json.Write(w, http.StatusNotFound, nil)
		return
	}

	json.Write(w, http.StatusOK, c)
}
