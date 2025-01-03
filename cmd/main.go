package main

import (
	"log"
	"rest-api-todo"
	"rest-api-todo/pkg/handler"
	"rest-api-todo/pkg/repository"
	"rest-api-todo/pkg/service"

	"github.com/spf13/viper"
)

func main() {
	if error := initConfig(); error != nil {
		log.Fatalf("error initializing configs: %s", error.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapitodo.Server)

	if error := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); error != nil {
		log.Fatalf("error occured while running http server: %s", error.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
