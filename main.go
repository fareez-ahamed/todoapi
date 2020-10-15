package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fareez-ahamed/todoapi/api"
	"github.com/fareez-ahamed/todoapi/store/sqlite"
)

func main() {
	store := sqlite.NewStore("test.db")
	todoHandler := &api.TodoHandler{
		Store: store,
	}
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: todoHandler,
	}

	// Listen to terminate signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("server closed: %v", err)
		}
	}()

	<-done
	log.Println("Shutting down server")

	err := httpServer.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("unable to shuttown error: %v", err)
	}
}
