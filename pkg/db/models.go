package db

import (
	"time"
)

type Balances struct {
	ID int64 `db:"id"`
	UserID int64 `db:"user_id"`
	Currency string `db:"currency"`
	Amount float64 `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type User struct {
	ID int64 `db:"id"`
	Username string `db:"username"`
	PasswordHash string `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TransactionHistory struct {
	ID int64 `db:"id"`
	UserID int64 `db:"user_id"`
	Asset string `db:"asset"`
	Amount float64 `db:"amount"`
	TransType string `db:"trans_type"`
	Status string `db:"status"`
	TransHash string `db:"trans_hash"`
	FromUserID *int64 `db:"from_user_id"`
	ToUserID *int64 `db:"to_user_id"`
	CreatedAt time.Time `db:"created_at"`
}