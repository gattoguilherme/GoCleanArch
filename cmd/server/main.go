package main

import (
	"GoCleanArch/configs"
	"GoCleanArch/internal/domain/entity"
	"GoCleanArch/internal/domain/repository"
	"GoCleanArch/internal/infra/database"
	"GoCleanArch/internal/infra/handler"
	"GoCleanArch/internal/infra/messaging"
	"GoCleanArch/internal/usecase"
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configuration
	configPath := flag.String("config", "./configs/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := configs.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	var orderRepo repository.OrderRepository
	var orderMessageQueue repository.OrderMessageQueue

	if cfg.Env == "dev" {
		log.Println("Running in development mode")
		// Mocks for dev environment
		orderRepoMock := database.NewOrderRepositoryMock()
		orderMessageQueue = messaging.NewOrderMessageQueueMock()

		// Pre-populating the mock database for the GET endpoint
		prePopulatedOrder := &entity.Order{
			ID:        "123",
			Data:      "Sample Order Data",
			OrderID:   123,
			Status:    "Completed",
			Paid:      true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		orderRepoMock.Save(prePopulatedOrder)
		orderRepo = orderRepoMock
	} else {
		log.Println("Running in production mode")
		// Real implementations for prod environment

		// SQS
		awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(cfg.Prod.AWS.Region))
		if err != nil {
			log.Fatalf("unable to load AWS config, %v", err)
		}
		sqsClient := sqs.NewFromConfig(awsCfg)
		orderMessageQueue = messaging.NewOrderMessageQueueSQS(sqsClient, cfg.Prod.AWS.SQSQueueURL)

		// MySQL
		db, err := sql.Open(cfg.Prod.DB.Driver, cfg.Prod.DB.DSN)
		if err != nil {
			log.Fatalf("could not connect to database: %v", err)
		}
		defer db.Close()
		orderRepo = database.NewOrderRepository(db)
	}

	// Use Cases
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderMessageQueue)
	getOrderUseCase := usecase.NewGetOrderByIDUseCase(orderRepo)

	// Handlers
	orderHandler := handler.NewOrderHandler(createOrderUseCase, getOrderUseCase)

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Logger) // Add a logger middleware
	r.Post("/orders", orderHandler.CreateOrder)
	r.Get("/orders/{orderId}", orderHandler.GetOrder)
	r.Get("/orders", orderHandler.GetAllOrders)

	log.Printf("Server is running on port %s", cfg.Server.Port)
	if err := http.ListenAndServe(cfg.Server.Port, r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
