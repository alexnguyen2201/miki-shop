// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type MtProductOrder struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	OrderID   int64 `json:"order_id"`
	// must be positive
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type Order struct {
	ID        int64     `json:"id"`
	OrderNo   string    `json:"order_no"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Size      sql.NullInt64  `json:"size"`
	Color     sql.NullString `json:"color"`
	Price     int64          `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
}

type User struct {
	ID                int64          `json:"id"`
	Username          string         `json:"username"`
	FullName          string         `json:"full_name"`
	Gender            sql.NullString `json:"gender"`
	DateOfBirth       sql.NullTime   `json:"date_of_birth"`
	Email             string         `json:"email"`
	HashedPassword    string         `json:"hashed_password"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt         time.Time      `json:"created_at"`
	PhoneNumber       sql.NullString `json:"phone_number"`
}
