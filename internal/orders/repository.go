package orders

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository для работы с пользователями
type Repository struct {
	db *pgxpool.Pool
}

// NewRepository создает новый репозиторий
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) New(ctx context.Context, order Order) (order_id int64, err error) {
	err = order.Validate()
	if err != nil {
		return -1, fmt.Errorf("ошибка валидации: %v", err)
	}

	var inserted_order Order
	err = r.db.QueryRow(ctx,
		`INSERT INTO orders (user_id, base_asset, quote_asset, type, side, price, amount, filled_amount, status, created_at, updated_at, fee) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
		RETURNING id;`,
		order.UserID, order.BaseAsset, order.QuoteAsset, order.Type, order.Side, order.Price, order.Amount, order.FilledAmount, order.Status, order.CreatedAt, order.UpdatedAt, order.Fee,
	).Scan(&inserted_order.ID)
	if err != nil {
		return -1, fmt.Errorf("ошибка при сохранении ордера: %v", err)
	}

	return 1, nil
}

// func (r *Repository) Get(order Order) ([]Order, error) {
// 	result := r.db.Find(&order, 10)
// }
