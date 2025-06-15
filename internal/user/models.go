package user

import "time"

// User - модель пользователя
type User struct {
	ID int64 `db:"id"`
	Username string `db:"username"`
	PasswordHash string `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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
