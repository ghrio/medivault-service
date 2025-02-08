-- name: GetAllRoles :many
SELECT id, role_name, created_at, updated_at, deleted_at
FROM roles;

-- name: GetActiveRoles :many
SELECT id, role_name, created_at, updated_at
FROM roles
WHERE deleted_at IS NULL;

-- name: CreateRole :one
INSERT INTO roles (role_name, created_at, updated_at, deleted_at)
VALUES ($1, NOW(), NOW(), NULL)
RETURNING id, role_name, created_at, updated_at, deleted_at;

-- name : UpdateRole :one
UPDATE roles
SET role_name = $2, updated_at = NOW()
WHERE id = $1
