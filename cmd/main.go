package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fojnk/todo-app3/internal/models"
	"github.com/fojnk/todo-app3/internal/repository"
	"github.com/fojnk/todo-app3/internal/service"
	"github.com/fojnk/todo-app3/internal/transport"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           Todo App Api
// @version         1.0
// @description     This is a simple version of todo app.
// @host      localhost:8000
// @BasePath  /
// @securitydefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("initConfig failed: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Load .env files failed: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Connection to Postgres DB failed: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := transport.NewHandler(services)

	srv := new(models.ServerApi)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("srv run failed: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
