package clients_test

import (
	"testing"

	"github.com/k4rldoherty/brige-backend/src/internal/clients"
)

func TestCreateClientDTOValidateInput(t *testing.T) {
	tests := []struct {
		name       string
		clientName string
		email      string
		joinCode   string
		logoURL    string
		wantErr    bool
		wantMsg    string
	}{
		{"valid all fields", "test", "test@test.com", "test", "www.test.com", false, ""},
		{"valid no logo url", "test", "test@test.com", "test", "", false, ""},
		{"invalid no name", "", "test@test.com", "test", "", false, ""},
		{"invalid no email", "test", "", "test", "", true, "email is required"},
		{"invalid no join code", "test", "test@test.com", "", "", true, "join_code is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clients.CreateClientDTO{Name: tt.name, Email: tt.email, JoinCode: tt.joinCode}
			err := c.ValidateInput()

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInput() error = %v, wantErr %v", err.Message, tt.wantErr)
			}
			if err != nil && err.Message != tt.wantMsg {
				t.Errorf("wanted message = %v, got %v", tt.wantMsg, err.Message)
			}
		})
	}
}

func TestUpdateClientDTOValidateInput(t *testing.T) {
	tests := []struct {
		name       string
		id         int32
		clientName string
		email      string
		logoURL    string
		wantErr    bool
		wantMsg    string
	}{
		{"valid all fields", 1, "test", "test@test.com", "www.test.com", false, ""},
		{"valid no logo url", 2, "test", "test@test.com", "", false, ""},
		{"invalid invalid id", 0, "test", "test@test.com", "", true, "id is required and must be a valid number greater than 0, and inside the int32 range"},
		{"invalid no email", 3, "test", "", "test", true, "email is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clients.UpdateClientDTO{ID: tt.id, Name: tt.clientName, Email: tt.email, LogoURL: tt.logoURL}
			err := c.ValidateInput()

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInput() error = %v, wantErr %v", err.Message, tt.wantErr)
			}
			if err != nil && err.Message != tt.wantMsg {
				t.Errorf("wanted message = %v, got %v", tt.wantMsg, err.Message)
			}
		})
	}
}
