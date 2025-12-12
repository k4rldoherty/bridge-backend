-- name: AddUser :one
INSERT INTO users (client_id, role_id, name, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUsersByClientID :many
SELECT u.name, u.email, c.name FROM users u
JOIN clients c ON u.client_id = c.id
WHERE client_id = $1;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING name;
