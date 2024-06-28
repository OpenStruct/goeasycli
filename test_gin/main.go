package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"test_gin/config"
	"test_gin/database"
	"test_gin/models"
	"test_gin/routers"
)

func main() {

	//config.LoadEnv()

	db := database.Init()

	migration := database.Migrations{
		DB: db,
		Models: []interface{}{
			&models.GoEasyCLITestUser{},
		},
	}

	database.RunMigrations(migration)

	port := config.CFG.V.GetString("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: routers.NewRouter(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("Server exiting")
}
