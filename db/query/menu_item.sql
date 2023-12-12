-- name: CreateMenu_item :one
INSERT INTO menu_item (
    title, category, price
) VALUES (
             $1, $2, $3
         )
    RETURNING *;

-- name: GetMenu_item :one
SELECT * FROM menu_item
WHERE menu_id = $1 LIMIT 1;

-- name: ListMenu_item :many
SELECT * FROM menu_item
ORDER BY menu_id
    LIMIT $1
OFFSET $2;

-- name: DeleteMenu_item :exec
DELETE FROM menu_item
WHERE menu_id = $1;