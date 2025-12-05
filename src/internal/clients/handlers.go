// Package clients - Handlers
package clients

import (
	"io"
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
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err, "failed to get clients", "handlers.GetClients")
		return
	}

	utils.Write(w, http.StatusOK, c)
}

func (h *handler) AddClient(w http.ResponseWriter, r *http.Request) {
	// Read in the body to a buffer
	d, err := io.ReadAll(r.Body)
	defer utils.CloseRequestBody(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err, "failed to parse body", "handlers.AddClient")
		return
	}

	// Call the service to add the client
	c, err := h.service.AddClient(r.Context(), d)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err, "failed to add client", "handlers.AddClient")
		return
	}

	utils.Write(w, http.StatusCreated, c)
}

func (h *handler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	defer utils.CloseRequestBody(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err, "failed to parse body", "handlers.UpdateClient")
		return
	}
	c, err := h.service.UpdateClient(r.Context(), d)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err, "failed to update client", "handlers.UpdateClient")
		return
	}

	utils.Write(w, http.StatusOK, c)
}
