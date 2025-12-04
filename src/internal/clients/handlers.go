// Package clients - Handlers
package clients

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/k4rldoherty/brige-backend/src/internal/utils"
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
		slog.Error("failed to get clients", "error", err, "location", "handlers.AddClient")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(c) == 0 || c == nil {
		utils.Write(w, http.StatusNotFound, nil)
		return
	}

	utils.Write(w, http.StatusOK, c)
}

func (h *handler) AddClient(w http.ResponseWriter, r *http.Request) {
	// Read in the body to a buffer
	d, err := io.ReadAll(r.Body)
	defer utils.CloseRequestBody(r)
	if err != nil {
		slog.Error("failed to parse body", "error", err, "location", "handlers.AddClient")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Call the service to add the client
	c, err := h.service.AddClient(r.Context(), d)
	if err != nil {
		slog.Error("failed to add client", "error", err, "location", "handlers.AddClient")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Write(w, http.StatusOK, c)
}
