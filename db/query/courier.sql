-- name: CreateCourier :one
INSERT INTO courier (
    title, phone
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: GetCourier :one
SELECT * FROM courier
WHERE courier_id = $1 LIMIT 1;

-- name: ListCourier :many
SELECT * FROM courier
ORDER BY courier_id
    LIMIT $1
OFFSET $2;

-- name: DeleteCourier :exec
DELETE FROM courier
WHERE courier_id = $1;