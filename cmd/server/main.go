package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/config"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/db/sqlc"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/internal/handler"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/internal/logger"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/internal/repository"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/internal/routes"
)

func main() {
	// Initialize logger
	if err := logger.Init(); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Log.Sync()

	logger.Log.Info("starting user service application...")

	// Connect to database
	db, err := config.ConnectDB()
	if err != nil {
		logger.Log.Fatal("failed to connect to database: " + err.Error())
	}
	defer db.Close()
	logger.Log.Info("connected to database successfully")

	// Initialize repository and handler
	queries := sqlc.New(db)
	repo := repository.NewUserRepository(queries)
	userHandler := handler.NewUserHandler(repo)

	// Initialize Fiber app
	app := fiber.New()

	// Register routes
	routes.Setup(app, userHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	logger.Log.Info("listening on port " + port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatal("failed to start server: " + err.Error())
	}
}
