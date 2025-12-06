package main

import (
	"log"

	"github.com/MerBerd/blog-app/internal/app/http/server"
	"github.com/MerBerd/blog-app/internal/handlers"
	"github.com/MerBerd/blog-app/internal/repositories"
	"github.com/MerBerd/blog-app/internal/services"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to load configs: %s", err.Error())
	}

	repos := repositories.NewRepository()
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := &server.Server{}
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to run server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
