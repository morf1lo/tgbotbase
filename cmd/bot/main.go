package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initializing all configs, repositories and services

	// logger, _ := zap.NewProduction()
	// states := &model.States{}
	// repos := repository.New(db, rdb)
	// services := service.New(repos)
	// handlers := handler.New(services)
	// bot := handler.NewBot(logger, handlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
