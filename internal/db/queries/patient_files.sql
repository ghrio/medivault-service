-- name: GetFilesByPatientCollection :many
SELECT id, patient_collection_id, link, created_at, deleted_at
FROM patent_files
WHERE patient_collection_id = $1 AND deleted_at IS NULL;

-- name: UploadPatentFile :one
INSERT INTO patent_files (patient_collection_id, link, created_at, deleted_at)
VALUES ($1, $2, NOW(), NULL)
RETURNING id, patient_collection_id, link, created_at, deleted_at;

-- name: SoftDeletePatentFile :exec
UPDATE patent_files
SET deleted_at = NOW()
WHERE id = $1;
