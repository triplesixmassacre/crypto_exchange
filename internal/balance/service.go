package balance

import (
	"context"
	"fmt"
	"log"
)

// Service для работы с балансами
type Service struct {
	repo *Repository
}

// NewService создает новый экземпляр сервиса
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetBalance возвращает баланс пользователя
func (s *Service) GetBalance(ctx context.Context, userID int64, asset string) (float64, error) {
	log.Printf("GetBalance вызван для user_id: %d, asset: %s", userID, asset)

	amount, err := s.repo.GetBalance(ctx, userID, asset)
	if err != nil {
		return 0, fmt.Errorf("ошибка получения баланса: %w", err)
	}

	return amount, nil
}

// UpdateBalance обновляет баланс пользователя
func (s *Service) UpdateBalance(ctx context.Context, userID int64, asset string, amount float64) (float64, error) {
	log.Printf("UpdateBalance вызван для user_id: %d, asset: %s, amount: %f",
		userID, asset, amount)

	// Проверяем существование баланса
	exists, err := s.repo.BalanceExists(ctx, userID, asset)
	if err != nil {
		return 0, fmt.Errorf("ошибка проверки существования баланса: %w", err)
	}
	if !exists {
		return 0, fmt.Errorf("баланс не найден для user_id: %d, asset: %s", userID, asset)
	}

	err = s.repo.UpdateBalance(ctx, userID, asset, amount)
	if err != nil {
		return 0, fmt.Errorf("ошибка обновления баланса: %w", err)
	}

	// Получаем обновленный баланс
	newAmount, err := s.repo.GetBalance(ctx, userID, asset)
	if err != nil {
		return 0, fmt.Errorf("ошибка получения обновленного баланса: %w", err)
	}

	return newAmount, nil
}
