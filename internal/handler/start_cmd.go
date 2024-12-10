package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) StartCMD(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	h.SendSimpleMessage(tgbot, update.Message.Chat.ID, "Hello!")
}
