// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (user_name, email, password_hash, bio)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateUserParams struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"password_hash"`
	Bio          string `json:"bio"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserName,
		arg.Email,
		arg.PasswordHash,
		arg.Bio,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
  id,
  user_name,
  email,
  password_hash,
  bio,
  updated_at,
  created_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Email,
		&i.PasswordHash,
		&i.Bio,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT
  id,
  user_name,
  email,
  password_hash,
  bio,
  updated_at,
  created_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Email,
		&i.PasswordHash,
		&i.Bio,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
