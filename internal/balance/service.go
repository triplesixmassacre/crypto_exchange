package balance

import (
	"context"
	"fmt"
	"log"

	pb "crypto_exchange/api/pb"
)

// Service реализует интерфейс BalanceServiceServer
type Service struct {
	pb.UnimplementedBalanceServiceServer
	repo *Repository
}

// NewService создает новый экземпляр сервиса
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetBalance возвращает баланс пользователя
func (s *Service) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	log.Printf("GetBalance вызван для user_id: %d, asset: %s", req.UserId, req.Asset)

	amount, err := s.repo.GetBalance(ctx, req.UserId, req.Asset)
	if err != nil {
		return nil, err
	}

	return &pb.GetBalanceResponse{ // Получить баланс из базы данных
		Amount: amount,
	}, nil
}

// UpdateBalance обновляет баланс пользователя
func (s *Service) UpdateBalance(ctx context.Context, req *pb.UpdateBalanceRequest) (*pb.UpdateBalanceResponse, error) {
	log.Printf("UpdateBalance вызван для user_id: %d, asset: %s, amount: %f",
		req.UserId, req.Asset, req.Amount)

	// Проверяем существование баланса
	exists, err := s.repo.BalanceExists(ctx, req.UserId, req.Asset)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("баланс не найден для user_id: %d, asset: %s", req.UserId, req.Asset)
	}

	err = s.repo.UpdateBalance(ctx, req.UserId, req.Asset, req.Amount)
	if err != nil {
		return nil, err
	}

	// Получаем обновленный баланс
	newAmount, err := s.repo.GetBalance(ctx, req.UserId, req.Asset)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateBalanceResponse{
		NewAmount: newAmount,
	}, nil
}
