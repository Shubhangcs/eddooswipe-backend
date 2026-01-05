package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/config"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/routes"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	// Loading all the required env
	cfg, err := config.Load()
	if err != nil {
		log.Println("failed to load env no file")
	}
	log.Println(".env loaded successfully")

	// Creating a database connection pool
	db, err := database.NewDatabaseConnection(database.Config{
		DatabaseURL: cfg.Database.DatabaseURL,
	})
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("connected to database")

	// Creating a router and initilizing all routes
	router, err := routes.NewRouter(routes.Config{
		Database:     db,
		JWTSecretKey: cfg.JWT.JWTSecretKey,
		JWTExpiry:    cfg.JWT.JWTExpiry,
		JWTSecretKeyPaySprint: cfg.JWT.JWTSecretKeyPaySprint,
	})
	if err != nil {
		return err
	}
	log.Println("routes initilized successfully")
	return start(router.EchoRouter, cfg.Server.ServerPort)
}

func start(router *echo.Echo, serverPort string) error {
	// Start Server
	go func() {
		if err := router.Start(serverPort); err != nil && err != http.ErrServerClosed {
			log.Println("server error:", err)
		}
	}()

	// Waiting for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutting the server down
	if err := router.Shutdown(ctx); err != nil {
		return fmt.Errorf("forced shutdown")
	}
	return nil
}
