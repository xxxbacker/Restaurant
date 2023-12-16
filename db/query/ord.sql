-- name: CreateOrd :one
INSERT INTO ord (
    account_id, courier_id, created_at
) VALUES (
             $1, $2, $3
         )
    RETURNING *;

-- name: GetOrd :one
SELECT * FROM ord
WHERE ord_id = $1 LIMIT 1;

-- name: ListOrd :many
SELECT * FROM ord
ORDER BY ord_id desc
LIMIT $1;

-- name: DeleteOrd :exec
DELETE FROM ord
WHERE ord_id = $1;