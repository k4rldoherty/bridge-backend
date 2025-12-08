package clients

import "errors"

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

func (c CreateClientDTO) ValidateInput() error {
	if c.Name == "" {
		return errors.New("400 Bad request : name is required")
	}
	if c.Email == "" {
		return errors.New("400 Bad request : email is required")
	}
	if c.JoinCode == "" {
		return errors.New("400 Bad request : join_code is required")
	}
	return nil
}

func (c UpdateClientDTO) ValidateInput() error {
	if c.ID == 0 {
		return errors.New("400 Bad request : id is required")
	}
	if c.Email == "" {
		return errors.New("400 Bad request : email is required")
	}
	if c.Name == "" {
		return errors.New("400 Bad request : name is required")
	}
	return nil
}
