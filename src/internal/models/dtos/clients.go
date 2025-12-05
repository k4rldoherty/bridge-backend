// Package dtos - Data Transfer Objects
package dtos

type CreateClientDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	JoinCode string `json:"join_code"`
	LogoURL  string `json:"logo_url"`
}

type UpdateClientDTO struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}
