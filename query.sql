-- name: CreateAccount :one
INSERT INTO accounts (
  id,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: CountAccounts :one
SELECT count(*) FROM accounts;

-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec 
DELETE FROM accounts
WHERE id = $1;