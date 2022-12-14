package main

import (
	"database/sql"
	"fmt"
	"github.com/Diulia/codebank-lessons/infrastructure/grpc/server"
	"github.com/Diulia/codebank-lessons/infrastructure/kafka"
	"github.com/Diulia/codebank-lessons/infrastructure/repository"
	"github.com/Diulia/codebank-lessons/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db := setupDb()
	defer db.Close()
	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	serveGrpc(processTransactionUseCase)
	cc := domain.NewCreditCard()
	cc.Number = "2222"
	cc.Name = "Diulia"
	cc.ExpirationYear = 2026
	cc.ExpirationMonth = 10
	cc.CVV = 987
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
	fmt.Println(err)
	}

}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))
	return producer
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("não conecta no db")
	}
	return db
}

func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	fmt.Println("Tá rodando certo até aqui")
	grpcServer.Serve()
}




// cc := domain.NewCreditCard()
// cc.Number = "2222"
// cc.Name = "Diulia"
// cc.ExpirationYear = 2026
// cc.ExpirationMonth = 10
// cc.CVV = 987
// cc.Limit = 1000
// cc.Balance = 0

// repo := repository.NewTransactionRepositoryDb(db)
// err := repo.CreateCreditCard(*cc)
// if err != nil {
// fmt.Println(err)
// }