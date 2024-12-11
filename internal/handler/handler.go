package handler

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	services struct{}       // Your services from 'internal/service'
	admins   map[int64]bool // Telegram bot admin ids
}

func New(logger *zap.Logger, services struct{}) *Handler {
	// There are 2 admins for example
	adminID0, _ := strconv.Atoi(os.Getenv("ADMIN_ID_0"))
	adminID1, _ := strconv.Atoi(os.Getenv("ADMIN_ID_1"))

	return &Handler{
		logger: logger,
		services: services,
		admins:   map[int64]bool{
			int64(adminID0): true,
			int64(adminID1): true,
		},
	}
}

// Filtering admins
func (h *Handler) HasPermission(update tgbotapi.Update) bool {
	return !update.Message.From.IsBot && h.admins[update.Message.From.ID]
}

// Just for saving 1 line
func (h *Handler) SendSimpleMessage(tgbot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	tgbot.Send(msg)
}

// Too
func (h *Handler) SendRemoveKeyboardMessage(tgbot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	tgbot.Send(msg)
}
