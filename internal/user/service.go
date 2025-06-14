package user

import (
	"context"
	"fmt"
)

// Service представляет сервис для работы с пользователями
type Service struct {
	repo *Repository
}

// NewService создает новый сервис
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Register регистрирует нового пользователя
func (s *Service) Register(ctx context.Context, req CreateUserRequest) (*UserResponse, error) {
	// Валидация входных данных
	if err := validateCreateRequest(req); err != nil {
		return nil, fmt.Errorf("не валидные данные: %v", err)
	}

	// Создание пользователя
	user, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// GetByID получает пользователя по ID
func (s *Service) GetByID(ctx context.Context, id int64) (*UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Update обновляет данные пользователя
func (s *Service) Update(ctx context.Context, id int64, req UpdateUserRequest) (*UserResponse, error) {
	// Валидация входных данных
	if err := validateUpdateRequest(req); err != nil {
		return nil, fmt.Errorf("не валидные данные: %v", err)
	}

	// Обновление пользователя
	user, err := s.repo.Update(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Delete удаляет пользователя
func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

// Login выполняет вход пользователя
func (s *Service) Login(ctx context.Context, username, password string) (*UserResponse, error) {
	user, err := s.repo.VerifyPassword(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Вспомогательные функции валидации
func validateCreateRequest(req CreateUserRequest) error {
	if len(req.Username) < 3 || len(req.Username) > 50 {
		return fmt.Errorf("имя пользователя должно быть от 3 до 50 символов")
	}
	if len(req.Password) < 8 {
		return fmt.Errorf("пароль должен быть не менее 8 символов")
	}
	return nil
}

func validateUpdateRequest(req UpdateUserRequest) error {
	if req.Username != "" && (len(req.Username) < 3 || len(req.Username) > 50) {
		return fmt.Errorf("имя пользователя должно быть от 3 до 50 символов")
	}
	if req.Password != "" && len(req.Password) < 8 {
		return fmt.Errorf("пароль должен быть не менее 8 символов")
	}
	return nil
}
