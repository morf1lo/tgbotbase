package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Handling commands
func (h *Handler) HandleCommand(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	switch update.Message.Command() {
	case "start": {
		h.StartCMD(tgbot, update)
		break
	}
	default: {
		h.SendSimpleMessage(tgbot, chatID, "Unknown command.")
	}
	}
}
