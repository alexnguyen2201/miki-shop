// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING id, username, full_name, gender, date_of_birth, email, hashed_password, password_changed_at, phone_number, created_at, updated_at, created_by, updated_by
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Gender,
		&i.DateOfBirth,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, full_name, gender, date_of_birth, email, hashed_password, password_changed_at, phone_number, created_at, updated_at, created_by, updated_by FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Gender,
		&i.DateOfBirth,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const getUserByUserName = `-- name: GetUserByUserName :one
SELECT id, username, full_name, gender, date_of_birth, email, hashed_password, password_changed_at, phone_number, created_at, updated_at, created_by, updated_by FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUserName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUserName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Gender,
		&i.DateOfBirth,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}
