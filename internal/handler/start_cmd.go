package handler

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/morf1lo/tgbotbase/internal/config"
	"github.com/morf1lo/tgbotbase/internal/localization"
)

func (h *Handler) StartCMD(b *gotgbot.Bot, ctx *ext.Context) error {
	telegramID := ctx.Message.From.Id

	b.SendMessage(
		telegramID,
		localization.GetMessage(config.ENGLISH_LANG, "startCMD.startMessage"),
		nil,
	)

	return handlers.NextConversationState(START_STATE)
}
