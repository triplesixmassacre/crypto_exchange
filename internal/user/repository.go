package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// Repository для работы с пользователями
type Repository struct {
	db *pgxpool.Pool
}

// NewRepository создает новый репозиторий
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

// Create создает нового пользователя
func (r *Repository) Create(ctx context.Context, req CreateUserRequest) (*User, error) {
	// Проверяем существует ли пользователь
	exists, err := r.Exists(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки существования пользователя: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("ошибка пользователь уже существует: %w", err)
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка хеширования пароля: %w", err)
	}

	// Создаем пользователя
	var user User
	err = r.db.QueryRow(ctx,
		`INSERT INTO users (username, password_hash, created_at, updated_at)
         VALUES ($1, $2, $3, $4)
         RETURNING id, username, password_hash, created_at, updated_at`,
		req.Username, string(hashedPassword), time.Now(), time.Now(),
	).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	return &user, nil
}

// GetByID получает пользователя по ID
func (r *Repository) GetByID(ctx context.Context, id int64) (*User, error) {
	var user User
	err := r.db.QueryRow(ctx,
		`SELECT id, username, password_hash, created_at, updated_at
         FROM users WHERE id = $1`,
		id,
	).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("ошибка пользователь не найден: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return &user, nil
}

// GetByUsername получает пользователя по имени пользователя
func (r *Repository) GetByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.db.QueryRow(ctx,
		`SELECT id, username, password_hash, created_at, updated_at
         FROM users WHERE username = $1`,
		username,
	).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("ошибка пользователь не найден: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return &user, nil
}

// Update обновляет данные пользователя
func (r *Repository) Update(ctx context.Context, id int64, req UpdateUserRequest) (*User, error) {
	// Получаем текущие данные пользователя
	user, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Обновляем поля, если они предоставлены
	if req.Username != "" {
		// Проверяем, не занято ли новое имя пользователя
		if req.Username != user.Username { // user с типом *User мы получили из метода GetByID (я сам охуел, что так можно, случайно получилось)
			exists, err := r.Exists(ctx, req.Username)
			if err != nil {
				return nil, fmt.Errorf("ошибка проверки существования пользователя: %w", err)
			}
			if exists {
				return nil, fmt.Errorf("ошибка пользователь уже существует: %w", err)
			}
		}
		user.Username = req.Username
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("ошибка хеширования пароля: %w", err)
		}
		user.PasswordHash = string(hashedPassword)
	}

	// Обновляем пользователя в БД
	err = r.db.QueryRow(ctx,
		`UPDATE users 
         SET username = $1, password_hash = $2, updated_at = $3
         WHERE id = $4
         RETURNING id, username, password_hash, created_at, updated_at`,
		user.Username, user.PasswordHash, time.Now(), id,
	).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("ошибка обновления пользователя: %w", err)
	}

	return user, nil
}

// Delete удаляет пользователя
func (r *Repository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.Exec(ctx,
		`DELETE FROM users WHERE id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("ошибка удаления пользователя: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("ошибка пользователь не найден: %w", err)
	}

	return nil
}

// Exists проверяет существование пользователя по имени
func (r *Repository) Exists(ctx context.Context, username string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`,
		username,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка проверки существования пользователя: %w", err)
	}
	return exists, nil
}

// VerifyPassword проверяет пароль пользователя
func (r *Repository) VerifyPassword(ctx context.Context, username, password string) (*User, error) {
	user, err := r.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки хэша: %w", err)
	}

	return user, nil
}
