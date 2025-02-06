-- name: GetPatientCollectionByUser :one
SELECT id, user_id, created_at
FROM patient_collections
WHERE user_id = $1;

-- name: CreatePatientCollection :one
INSERT INTO patient_collections (user_id, created_at)
VALUES ($1, NOW())
RETURNING id, user_id, created_at;
