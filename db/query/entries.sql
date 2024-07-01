-- name: CreateEntrie :one
INSERT INTO entries (
  account_id, 
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetEntrie :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: Listentries :many
SELECT * FROM entries
ORDER BY account_id;

-- name: UpdateEntrie :exec
UPDATE entries
  set amount = $2
WHERE id = $1;

-- name: DeleteEntrie :exec
DELETE FROM entries
WHERE id = $1;