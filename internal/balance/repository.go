package balance

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository для работы с балансами
type Repository struct {
	db *pgxpool.Pool
}

// NewRepository создает новый репозиторий
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

// GetBalance получает баланс пользователя
func (r *Repository) GetBalance(ctx context.Context, userID int64, asset string) (float64, error) {
	var amount float64
	err := r.db.QueryRow(ctx, // QueryRow выполняет запрос и ожидает возврат, то есть amount
		`SELECT amount 
		FROM balances 
		WHERE user_id = $1 AND asset = $2`,
		userID, asset).Scan(&amount) // с помощью Scan копируем результат запроса в переменную amount
	if err != nil {
		return 0, fmt.Errorf("ошибка получения баланса: %v", err)
	}
	return amount, nil
}

// UpdateBalance обновляет баланс пользователя
func (r *Repository) UpdateBalance(ctx context.Context, userID int64, asset string, amount float64) error {
	_, err := r.db.Exec(ctx, // Exec выполняет запрос и не ожидает возврата
		`INSERT INTO balances (user_id, asset, amount) 
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, asset) 
		DO UPDATE SET amount = balances.amount + $3`,
		userID, asset, amount) // ON CONFLICT обновляет запись, если она уже есть
	if err != nil {
		return fmt.Errorf("ошибка обновления баланса: %v", err)
	}
	return nil
}

// CreateWallet создает новый кошелек
func (r *Repository) CreateWallet(ctx context.Context, userID int64, publicKey string, seedPhrase string) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO users (id, public_key, seed_phrase) 
		VALUES ($1, $2, $3)`,
		userID, publicKey, seedPhrase)
	if err != nil {
		return fmt.Errorf("ошибка создания кошелька: %v", err)
	}
	return nil
}
