package main

import (
	"barber/src/database"
	"barber/src/repositories"
	"barber/src/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var (
	err      error
	httpPort = ":3333"
)

func main() {
	repositories.Client, err = database.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	repositories.Database = repositories.Client.Database("gobarber")
	repositories.UserCollection = repositories.NewUserCollection()

	r := routes.LoadUserRoutes()

	server := &http.Server{
		Addr:    httpPort,
		Handler: r,
	}

	// start API server
	go func() {
		fmt.Printf("API listening on port %s\n", httpPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	// The context is used to inform the server it has 60 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("disconnecting",
		"component", "database",
	)
	err = repositories.Client.Disconnect(ctx)
	if err != nil {
		log.Printf("could not gracefully disconnect to database: %v\n", err)
		os.Exit(1)
	}
	log.Println("disconnected",
		"component", "database",
	)

	log.Println("server exiting")
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server forced to shutdown: %v\n", err)
		os.Exit(1)
	}
	log.Println("server successfully stopped")
}
