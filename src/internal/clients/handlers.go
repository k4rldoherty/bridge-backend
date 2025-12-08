// Package clients - Handlers
package clients

import (
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
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

// the handler is responsible for handling the request and returning a response
func (h *handler) GetClients(w http.ResponseWriter, r *http.Request) {
	c, err := h.service.GetClients(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(c) == 0 {
		utils.Write(w, http.StatusNoContent, nil)
	}

	utils.Write(w, http.StatusOK, c)
}

func (h *handler) AddClient(w http.ResponseWriter, r *http.Request) {
	// Read in the body to a buffer
	d, err := io.ReadAll(r.Body)
	defer utils.CloseRequestBody(r, h.logger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Call the service to add the client
	c, err := h.service.AddClient(r.Context(), d)
	if err != nil {
		if strings.Contains(err.Error(), "400") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Write(w, http.StatusCreated, c)
}

func (h *handler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	defer utils.CloseRequestBody(r, h.logger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c, err := h.service.UpdateClient(r.Context(), d)
	if err != nil {
		if strings.Contains(err.Error(), "400") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Write(w, http.StatusOK, c)
}

func (h *handler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "could not parse id from path", http.StatusInternalServerError)
		return
	}
	err := h.service.DeleteClient(r.Context(), []byte(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Write(w, http.StatusOK, nil)
}
