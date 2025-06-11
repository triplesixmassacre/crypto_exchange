package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Config содержит настройки подключения к БД
type Config struct {
	Host string
	Port int
	User string
	Password string
	DBName string
}

// NewPostgresDB создает новое подключение к PostgreSQL
func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", // dsn - data source name, строка для подключения к БД
	cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	
	pool, err := pgxpool.New(context.Background(), dsn) // pgxpool - пул соединений с БД
	if err != nil {
		log.Printf("ошибка подключения к БД: %v", err)
		return nil, err
	}

	// Проверяем подключение
	if err := pool.Ping(context.Background()); err != nil {
		log.Printf("ошибка пинга БД: %v", err)
		return nil, err
	}

	log.Println("Успешное подключение к PostgreSQL")
	return pool, nil
}
