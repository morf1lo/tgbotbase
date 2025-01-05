package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/morf1lo/tgbotbase/internal/config"
	"github.com/morf1lo/tgbotbase/internal/handler"
	"github.com/morf1lo/tgbotbase/internal/repository"
	"github.com/morf1lo/tgbotbase/internal/repository/postgres"
	"github.com/morf1lo/tgbotbase/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	// Initializing all configs, repositories and services
	logger, _ := zap.NewProduction()

	if err := initEnv(); err != nil {
		logger.Sugar().Fatalf("failed to load environment variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logger.Sugar().Fatalf("failed to initialize yaml config: %s", err.Error())
	}

	if err := config.LoadLocalizations(); err != nil {
		logger.Sugar().Fatalf("failed to load localizations: %s", err.Error())
	}

	postgresConfig := &config.PostgresConfig{
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		DBName: os.Getenv("POSTGRES_DBNAME"),
		SSLMode: os.Getenv("POSTGRES_SSLMODE"),
	}
	db, err := postgres.NewPostgresDatabase(ctx, postgresConfig)
	if err != nil {
		logger.Sugar().Fatalf("failed to connect to Postgres: %s", err.Error())
	}
	err = db.Ping(ctx)
	if err != nil {
		logger.Sugar().Fatalf("failed to connect to Postgres: %s", err)
	}
	logger.Sugar().Info("Successfully connected to Postgres")

	redisOptions := &redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
		Protocol: 3,
		ReadTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
	}
	rdb := redis.NewClient(redisOptions)
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Sugar().Fatalf("failed to ping Redis: %s", err)
	}
	logger.Sugar().Infof("Successfully connected to Redis: %s", pong)

	botConfig := &config.BotConfig{
		Token: os.Getenv("BOT_TOKEN"),
		Debug: viper.GetBool("bot.debug"),
	}

	repos := repository.New(db, rdb)
	services := service.New(repos)
	handlers := handler.New(ctx, logger, services, botConfig)

	logger.Info("Telegam Bot Started. Press Ctrl + C to exit.")

	handlers.RunBot()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func initEnv() error {
	return godotenv.Load()
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
