package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ctx = context.Background()

func main() {
	// Initializing all configs, repositories and services
	logger, _ := zap.NewProduction()

	if err := initEnv(); err != nil {
		logger.Sugar().Fatalf("failed to load environment variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logger.Sugar().Fatalf("failed to initialize yaml config: %s", err.Error())
	}

	// postgresConfig := &config.PostgresConfig{
	// 	Username: os.Getenv("POSTGRES_USERNAME"),
	// 	Password: os.Getenv("POSTGRES_PASSWORD"),
	// 	Host: os.Getenv("POSTGRES_HOST"),
	// 	Port: os.Getenv("POSTGRES_PORT"),
	// 	DBName: os.Getenv("POSTGRES_DATABASE"),
	// 	SSLMode: os.Getenv("POSTGRES_SSLMODE"),
	// }
	// db, err := db.NewPostgresDatabase(ctx, postgresConfig)
	// if err != nil {
	// 	logger.Sugar().Fatalf("failed to connect to PostgreSQL Database: %s", err.Error())
	// }

	// redisOptions := &redis.Options{
	// 	Addr: os.Getenv("REDIS_ADDR"),
	// 	Username: os.Getenv("REDIS_USERNAME"),
	// 	Password: os.Getenv("REDIS_PASSWORD"),
	// 	DB: 0,
	// 	Protocol: 2,
	// 	ReadTimeout: time.Second * 5,
	// 	WriteTimeout: time.Second * 5,
	// }
	// rdb := redis.NewClient(redisOptions)

	// repos := repository.New(db, rdb)
	// services := service.New(repos)
	// handlers := handler.New(services)
	// bot := handler.NewBot(logger, handlers)
	// bot.Start(os.Getenv("BOT_TOKEN", viper.GetBool("bot.debug")))

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
