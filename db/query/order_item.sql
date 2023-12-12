-- name: CreateOrder_item :one
INSERT INTO order_item (
    title, price, menu_id, ord_id
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING *;

-- name: GetOrder_item :one
SELECT * FROM order_item
WHERE order_item_id = $1 LIMIT 1;

-- name: ListOrder_item :many
SELECT * FROM order_item
ORDER BY order_item_id
LIMIT $1;

-- name: DeleteOrder_item :exec
DELETE FROM order_item
WHERE order_item_id = $1;