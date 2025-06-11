package balance

import (
	"context"
	"fmt"
	"log"

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
	log.Printf("Начало транзакции для user_id: %d, asset: %s, amount: %f", userID, asset, amount)

	// Начало транзакции
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("ошибка начала транзакции: %v", err) // юзаем Errorf, а не Println, потому что ждём возврата типа error
	}
	defer tx.Rollback(ctx)

	// Блокируем строку для обновления с помощью FOR UPDATE
	var currentAmount float64
	err = tx.QueryRow(ctx,
		`SELECT amount 
		FROM balances 
		WHERE user_id = $1 AND asset = $2
		FOR UPDATE`,
		userID, asset).Scan(&currentAmount)
	if err != nil {
		return fmt.Errorf("ошибка получения баланса: %v", err)
	}

	log.Printf("Текущий баланс: %f", currentAmount)

	// Проверяем баланс
	if currentAmount + amount < 0 { // бля я не ебу как проверить условие помимо 0, а типа вот суммы
		return fmt.Errorf("недостаточно средств")
	}

	// Обновление баланса
	_, err = tx.Exec(ctx,
		`UPDATE balances
	SET amount = amount + $3
	WHERE user_id = $1 AND asset = $2`,
		userID, asset, amount)
	if err != nil {
		return fmt.Errorf("ошибка обновления баланса: %v", err)
	}

	// Окончание транзакции
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("ошибка подтверждения транзакции: %v", err)
	}

	log.Printf("Транзакция успешно завершена для user_id: %d, asset: %s", userID, asset)
	return nil
}

// BalanceExists проверяет существование баланса
func (r *Repository) BalanceExists(ctx context.Context, userID int64, asset string) (bool, error) {
	var exists bool = true
	err := r.db.QueryRow(ctx,
		`SELECT EXISTS(
			SELECT 1 
			FROM balances 
			WHERE user_id = $1 AND asset = $2
		)`,
		userID, asset).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка проверки существования баланса: %v", err)
	}
	return exists, nil
}
