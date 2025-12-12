// Package clients - Handlers
package clients

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/k4rldoherty/brige-backend/src/internal/logger"
	"github.com/k4rldoherty/brige-backend/src/internal/utils"
)

func NewHandler(service Service, logger *logger.Logger) *handler {
	return &handler{service: service, logger: logger}
}

func (h *handler) GetClients(w http.ResponseWriter, r *http.Request) {
	c, err := h.service.GetClients(r.Context())
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	if len(c) == 0 {
		utils.Write(w, http.StatusNoContent, nil)
	}

	utils.Write(w, http.StatusOK, c)
}

func (h *handler) AddClient(w http.ResponseWriter, r *http.Request) {
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

	c, err := h.service.AddClient(r.Context(), d)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	utils.Write(w, http.StatusCreated, c)
}

func (h *handler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	d, e := io.ReadAll(r.Body)
	defer utils.CloseRequestBody(r, h.logger)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	c, err := h.service.UpdateClient(r.Context(), d)
	if err != nil {
		http.Error(w, err.Message, err.Status)
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

	err := h.service.DeleteClient(r.Context(), id)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	utils.Write(w, http.StatusOK, nil)
}
