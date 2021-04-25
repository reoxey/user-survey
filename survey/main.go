package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"survey/core"
	"survey/logger"
	"survey/repo/mongo"
	"survey/router"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	dbName := os.Getenv("DB_NAME")

	log := logger.New() // custom logger

	// context to timeout after 10s sent to mongo connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// initialise mongo repo
	// Mongo implementing Repo interface
	// Can be replaced by any DB
	repo, err := mongo.NewRepo(ctx, dsn, dbName)
	if err != nil {
		log.Fatalf("mongo: %s\n", err)
	}

	// Implements Service interface
	service := core.NewService(repo)

	r := router.New(log, true) // Gin router

	r.Handle(service)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	go func() { // started http server
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// channel listening for interrupts to ensure graceful shutdown of the http server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
