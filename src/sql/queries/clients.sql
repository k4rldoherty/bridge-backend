-- name: AddClient :one
INSERT INTO clients (name, email, join_code, logo_url) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetClients :many
SELECT * FROM clients;

-- name: UpdateClient :one
UPDATE clients SET name = $2, email = $3, logo_url = $4, updated_at = NOW() 
WHERE id = $1 
RETURNING *;

-- name: DeleteClient :one
DELETE FROM clients WHERE id = $1 RETURNING name;
