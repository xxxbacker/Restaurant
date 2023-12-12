-- name: CreateAccount :one
INSERT INTO account (
    post, first_name, last_name, password, email, phone, created_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         )
    RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE account_id = $1 LIMIT 1;

-- name: GetAccountForPassword :one
SELECT * FROM account
WHERE phone = $1 and password = $2
LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY account_id
    LIMIT $1
OFFSET $2;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE account_id = $1;

