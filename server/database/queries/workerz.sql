-- name: InsertWorkerz :one
INSERT INTO workerz (
    full_name,
    email,
    user_id,
    role,
    country,
    salary,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetWorkerzByUserID :one
SELECT * FROM workerz WHERE user_id = $1;

-- name: GetWorkerzByEmail :one
SELECT * FROM workerz WHERE email = $1;

-- name: GetAllWorkerz :many
SELECT * FROM workerz;

-- name: ListWorkerzByCountry :many
SELECT * FROM workerz
WHERE country = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateWorkerzEmail :one
UPDATE workerz SET email = $1 WHERE id = $2 RETURNING *;

-- name: DeleteWorkerz :exec
DELETE FROM workerz WHERE id = $1;

-- name: DeleteAllWorkerz :exec
DELETE FROM workerz;