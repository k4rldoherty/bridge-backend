// Package utils - A collection of utility functions that i keep using
package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/k4rldoherty/brige-backend/src/internal/logger"
)

// CloseRequestBody closes the request body, logging any errors
func CloseRequestBody(r *http.Request, l *logger.Logger) {
	err := r.Body.Close()
	if err != nil {
		l.Error("failed to close request body", "error", err, "location", "utils.CloseRequestBody")
	}
}

// Write writes a response to the http.ResponseWriter, setting headers used in most responses
func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
