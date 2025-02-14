-- name: GetAllUsers :many
SELECT id, name, email, password, created_at, updated_at, deleted_at
FROM users;

-- name: GetActiveUsers :many
SELECT id, name, email, password, created_at, updated_at, deleted_at
FROM users
WHERE deleted_at IS NULL;

-- name: CreateUser :one
INSERT INTO users (name, email, password, created_at, updated_at, deleted_at)
VALUES ($1, $2, $3, NOW(), NOW(), NULL)
RETURNING id, name, email, password, created_at, updated_at, deleted_at;

-- name: UpdateUserEmail :exec
UPDATE users
SET email = $1, updated_at = NOW()
WHERE id = $2 AND deleted_at IS NULL;

-- name: SoftDeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL;

-- name: CountUsers :one
SELECT COUNT(*) AS user_count
FROM users
WHERE deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT id, name, email, password, created_at, updated_at, deleted_at
FROM users
WHERE email = $1 AND deleted_at IS NULL;
RETURNING id, name, email, password, created_at, updated_at, deleted_at;
