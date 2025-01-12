// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"
)

const userLoginByEmail = `-- name: UserLoginByEmail :one
SELECT id, name, email, password, created_at FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) UserLoginByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, userLoginByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
