// Package utils - A collection of utility functions/types that i use across packages
package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/k4rldoherty/brige-backend/src/internal/logger"
)

type APIError struct {
	Status  int
	Message string
}

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

// ToNullString returns a null string object from a string, depending on if it is empty or not
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

func ToNullInt32(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{
			Valid: false,
		}
	}
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

// UnmarshalJSON turns a json object into a specified type, returning an APIError struct if an error occurs
func UnmarshalJSON(d []byte, v any) *APIError {
	err := json.Unmarshal(d, v)
	if err != nil {
		return &APIError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return nil
}
