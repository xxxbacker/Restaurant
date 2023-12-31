// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: order_item.sql

package db

import (
	"context"
	"database/sql"
)

const createOrder_item = `-- name: CreateOrder_item :one
INSERT INTO order_item (
    title, price, menu_id, ord_id
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING order_item_id, title, price, created_at, menu_id, ord_id
`

type CreateOrder_itemParams struct {
	Title  sql.NullString `json:"title"`
	Price  int32          `json:"price"`
	MenuID sql.NullInt64  `json:"menu_id"`
	OrdID  sql.NullInt64  `json:"ord_id"`
}

func (q *Queries) CreateOrder_item(ctx context.Context, arg CreateOrder_itemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrder_item,
		arg.Title,
		arg.Price,
		arg.MenuID,
		arg.OrdID,
	)
	var i OrderItem
	err := row.Scan(
		&i.OrderItemID,
		&i.Title,
		&i.Price,
		&i.CreatedAt,
		&i.MenuID,
		&i.OrdID,
	)
	return i, err
}

const deleteOrder_item = `-- name: DeleteOrder_item :exec
DELETE FROM order_item
WHERE order_item_id = $1
`

func (q *Queries) DeleteOrder_item(ctx context.Context, orderItemID int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrder_item, orderItemID)
	return err
}

const getOrder_item = `-- name: GetOrder_item :one
SELECT order_item_id, title, price, created_at, menu_id, ord_id FROM order_item
WHERE order_item_id = $1 LIMIT 1
`

func (q *Queries) GetOrder_item(ctx context.Context, orderItemID int32) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrder_item, orderItemID)
	var i OrderItem
	err := row.Scan(
		&i.OrderItemID,
		&i.Title,
		&i.Price,
		&i.CreatedAt,
		&i.MenuID,
		&i.OrdID,
	)
	return i, err
}

const listOrder_item = `-- name: ListOrder_item :many
SELECT order_item_id, title, price, created_at, menu_id, ord_id FROM order_item
ORDER BY order_item_id
LIMIT $1
`

func (q *Queries) ListOrder_item(ctx context.Context, limit int32) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, listOrder_item, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderItem{}
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.OrderItemID,
			&i.Title,
			&i.Price,
			&i.CreatedAt,
			&i.MenuID,
			&i.OrdID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
