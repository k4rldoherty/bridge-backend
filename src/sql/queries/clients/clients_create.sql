-- name: AddClient :one
INSERT INTO clients (name, email, join_code, logo_url) VALUES ($1, $2, $3, $4) RETURNING *;
