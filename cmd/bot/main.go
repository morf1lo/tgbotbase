package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// Initializing all configs, repositories and services
	logger, _ := zap.NewProduction()

	if err := initEnv(); err != nil {
		logger.Sugar().Fatalf("failed to load environment variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logger.Sugar().Fatalf("failed to initialize yaml config: %s", err.Error())
	}

	// states := &model.States{}
	// repos := repository.New(db, rdb)
	// services := service.New(repos)
	// handlers := handler.New(services)
	// bot := handler.NewBot(logger, handlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func initEnv() error {
	return godotenv.Load()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
