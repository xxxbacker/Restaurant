// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: ord.sql

package db

import (
	"context"
	"database/sql"
)

const createOrd = `-- name: CreateOrd :one
INSERT INTO ord (
    account_id, courier_id
) VALUES (
             $1, $2
         )
    RETURNING ord_id, ord_date, created_at, account_id, courier_id
`

type CreateOrdParams struct {
	AccountID sql.NullInt64 `json:"account_id"`
	CourierID sql.NullInt64 `json:"courier_id"`
}

func (q *Queries) CreateOrd(ctx context.Context, arg CreateOrdParams) (Ord, error) {
	row := q.db.QueryRowContext(ctx, createOrd, arg.AccountID, arg.CourierID)
	var i Ord
	err := row.Scan(
		&i.OrdID,
		&i.OrdDate,
		&i.CreatedAt,
		&i.AccountID,
		&i.CourierID,
	)
	return i, err
}

const deleteOrd = `-- name: DeleteOrd :exec
DELETE FROM ord
WHERE ord_id = $1
`

func (q *Queries) DeleteOrd(ctx context.Context, ordID int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrd, ordID)
	return err
}

const getOrd = `-- name: GetOrd :one
SELECT ord_id, ord_date, created_at, account_id, courier_id FROM ord
WHERE ord_id = $1 LIMIT 1
`

func (q *Queries) GetOrd(ctx context.Context, ordID int32) (Ord, error) {
	row := q.db.QueryRowContext(ctx, getOrd, ordID)
	var i Ord
	err := row.Scan(
		&i.OrdID,
		&i.OrdDate,
		&i.CreatedAt,
		&i.AccountID,
		&i.CourierID,
	)
	return i, err
}

const listOrd = `-- name: ListOrd :many
SELECT ord_id, ord_date, created_at, account_id, courier_id FROM ord
ORDER BY ord_id desc
LIMIT $1
`

func (q *Queries) ListOrd(ctx context.Context, limit int32) ([]Ord, error) {
	rows, err := q.db.QueryContext(ctx, listOrd, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ord{}
	for rows.Next() {
		var i Ord
		if err := rows.Scan(
			&i.OrdID,
			&i.OrdDate,
			&i.CreatedAt,
			&i.AccountID,
			&i.CourierID,
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