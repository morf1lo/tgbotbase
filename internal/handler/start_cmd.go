package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/morf1lo/tgbotbase/internal/config"
	"github.com/morf1lo/tgbotbase/internal/localization"
)

func (h *Handler) StartCMD(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	h.SendSimpleMessage(
		tgbot,
		update.Message.Chat.ID,
		localization.GetMessage(config.ENGLISH_LANG, "startCMD.message"),
	)
}
