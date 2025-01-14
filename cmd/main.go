package main

import (
	"os"
	restapitodo "rest-api-todo"
	"rest-api-todo/pkg/handler"
	"rest-api-todo/pkg/repository"
	"rest-api-todo/pkg/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	
	if error := initConfig(); error != nil {
		logrus.Fatalf("error initializing configs: %s", error.Error())
	}

	if error := gotenv.Load(); error != nil {
		logrus.Fatalf("error loading env variables: %s", error.Error())
	}

	db, error := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if error != nil {
		logrus.Fatalf("failed to initialize db: %s", error.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapitodo.Server)

	if error := srv.Run(viper.GetString("port"), handlers.InitRoutes()); error != nil {
		logrus.Fatalf("error occured while running http server: %s", error.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
