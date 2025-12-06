package main

import (
	"log"

	"github.com/MerBerd/blog-app/internal/app/http/server"
	"github.com/MerBerd/blog-app/internal/handlers"
	"github.com/MerBerd/blog-app/internal/repositories"
	"github.com/MerBerd/blog-app/internal/services"
)

func main() {
	repos := repositories.NewRepository()
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := &server.Server{}
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to run server: %s", err.Error())
	}
}
