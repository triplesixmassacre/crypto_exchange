package user

import "time"

// User - модель пользователя
type User struct {
	ID           int64
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CreateUserRequest - запрос на создание пользователя
type CreateUserRequest struct {
	Username string
	Password string
}

// UpdateUserRequest - запрос на обновление пользователя
type UpdateUserRequest struct {
	Username string
	Password string
}

// UserResponse - ответ с данными пользователя
type UserResponse struct {
	ID        int64
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
