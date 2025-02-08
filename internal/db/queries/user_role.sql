-- name: GetRoles :many
SELECT id, role_name
FROM roles
WHERE deleted_at IS NULL;

-- name: GetUsersWithRoles :many
SELECT users.id, users.name, users.email, roles.role_name
FROM users
JOIN user_roles ON users.id = user_roles.user_id
JOIN roles ON user_roles.role_id = roles.id
WHERE users.deleted_at IS NULL AND roles.deleted_at IS NULL;

-- name: AssignRoleToUser :exec
INSERT INTO user_roles (user_id, role_id, status)
VALUES ($1, $2, TRUE)
ON CONFLICT (user_id, role_id) DO UPDATE SET status = TRUE;

-- name: RemoveRoleFromUser :exec
DELETE FROM user_roles
WHERE user_id = $1 AND role_id = $2;
