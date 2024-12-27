package main

import (
	"log"
	"rest-api-todo"
	"rest-api-todo/pkg/handler"
	"rest-api-todo/pkg/repository"
	"rest-api-todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapitodo.Server)

	if error := srv.Run("8000", handlers.InitRoutes()); error != nil {
		log.Fatalf("error occured while running http server: %s", error.Error()	)
	}
}
