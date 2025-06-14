package main

import (
	"log"
	"net"

	"crypto_exchange/internal/balance"
	"crypto_exchange/pkg/db"
)

func main() {
	// Подключаемся к базе данных
	dbConfig := db.Config{
		Host: "localhost",
		Port: 5432,
		User: "postgres",
		Password: "postgres",
		DBName: "crypto_exchange",
	}
	pool, err := db.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer pool.Close()

	// Создаем репозиторий и сервис
	balanceRepo := balance.NewRepository(pool)
	balanceService := balance.NewService(balanceRepo) // баби буп надо будет использовать, когда сделаем обработчика подключения

	// Создаем TCP слушателя
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка прослушивания порта: %v", err)
	}
	defer listener.Close()

	log.Print("Сервер запущен на порту 50051")

}
