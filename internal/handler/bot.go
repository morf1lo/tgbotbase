package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	logger *zap.Logger
	handler *Handler
}

func NewBot(logger *zap.Logger, handler *Handler) *Bot {
	return &Bot{
		logger: logger,
		handler: handler,
	}
}

// Start bot
func (b *Bot) Start(token string, debug bool) {
	tgbot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		b.logger.Sugar().Fatalf("failed to start bot: %s", err.Error())
	}
	tgbot.Debug = debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	
	updates := tgbot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			go b.handler.HandleCommand(tgbot, update)
		} else if update.Message != nil && !update.Message.IsCommand() {
			go b.handler.HandleMessage(tgbot, update)
		} else if update.CallbackQuery != nil {
			go b.handler.CallbackQueryHandler(tgbot, update.CallbackQuery)
		}
	}
}
