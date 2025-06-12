package db

import (
	"time"
)

type Balances struct {
	UserID int64 `db:"user_id"`
	Asset string `db:"asset"`
	Amount float64 `db:"amount"`
}

type Users struct {
	ID int64 `db:"id"`
	PublicKey string `db:"public_key"`
	CreatedAt time.Time `db:"created_at"`
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