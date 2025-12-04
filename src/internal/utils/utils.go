// Package utils - A collection of utility functions that i keep using
package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func CloseRequestBody(r *http.Request) {
	err := r.Body.Close()
	if err != nil {
		slog.Error("failed to close request body", "error", err, "location", "utils.CloseRequestBody")
	}
}

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
