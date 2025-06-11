package main

import (
	"log"
	"net"

	"crypto_exchange/api/pb"
	"crypto_exchange/internal/balance"
	"crypto_exchange/pkg/db"

	"google.golang.org/grpc"
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
	defer pool.Close() // закрываем соединение с БД при выходе из функции

	// Создаем репозиторий и сервис
	balanceRepo := balance.NewRepository(pool)
	balanceService := balance.NewService(balanceRepo)

	// Создаем TCP слушателя
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка прослушивания порта: %v", err)
	}

	// Создаем новый gRPC сервер
	s := grpc.NewServer()

	// Регистрируем наш сервис
	pb.RegisterBalanceServiceServer(s, balanceService)

	log.Println("Сервер запущен на порту :50051")

	// Запускаем сервер
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
