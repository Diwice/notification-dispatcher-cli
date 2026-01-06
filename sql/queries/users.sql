-- name: CreateUser :one
INSERT INTO users(preferences, mail, phone_number)
VALUES(
	?,
	?,
	?
)
RETURNING *;
-- name: UpdateUserByID :one
UPDATE users
SET preferences = ?, mail = ?, phone_number = ?
WHERE id = ?
RETURNING *;
-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = ?;
-- name: DeleteAllUsers :exec
DELETE FROM users;
-- name: GetUserById :one
SELECT * FROM users
WHERE id = ?;
