-- name: CreateEnquiry :exec
INSERT INTO enquiries (name, phone, email, message)
VALUES ($1, $2, $3, $4)
    RETURNING id, created_at;

-- name: GetEnquiries :many
SELECT id, name, phone, email, message, created_at
FROM enquiries
ORDER BY created_at DESC;