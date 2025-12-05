-- name: UpdateClient :one
UPDATE clients SET name = $2, email = $3, logo_url = $4, updated_at = NOW() 
WHERE id = $1 
RETURNING *;
