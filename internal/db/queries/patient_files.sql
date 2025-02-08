-- name: GetFilesByPatientCollection :many
SELECT id, patient_collection_id, link, created_at, deleted_at
FROM patient_files
WHERE patient_collection_id = $1 AND deleted_at IS NULL;

-- name: UploadPatentFile :one
INSERT INTO patient_files (patient_collection_id, link, created_at, deleted_at)
VALUES ($1, $2, NOW(), NULL)
RETURNING id, patient_collection_id, link, created_at, deleted_at;

-- name: SoftDeletePatentFile :exec
UPDATE patient_files
SET deleted_at = NOW()
WHERE id = $1;
