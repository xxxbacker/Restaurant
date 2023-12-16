-- name: CreateCheque :one
INSERT INTO cheque (
    price, ord_id, created_at
) VALUES (
             $1, $2, $3
         )
    RETURNING *;

-- name: GetCheque :one
SELECT * FROM cheque
WHERE cheque_id = $1 LIMIT 1;

-- name: ListCheque :many
SELECT * FROM cheque
ORDER BY cheque_id DESC
LIMIT $1;

-- name: DeleteCheque :exec
DELETE FROM cheque
WHERE cheque_id = $1;