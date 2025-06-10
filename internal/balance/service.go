package balance

import (
	"context"
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

	err := s.repo.UpdateBalance(ctx, req.UserId, req.Asset, req.Amount)
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

// CreateWallet создает новый кошелек для пользователя
func (s *Service) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletResponse, error) {
	log.Printf("CreateWallet вызван для user_id: %d, public_key: %s",
		req.UserId, req.PublicKey)

	seedPhrase := "test seed phrase" 

	err := s.repo.CreateWallet(ctx, req.UserId, req.PublicKey, seedPhrase)
	if err != nil {
		return nil, err
	}

	return &pb.CreateWalletResponse{
		Success: true,
	}, nil
}
