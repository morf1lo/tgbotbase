package handler

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/morf1lo/tgbotbase/internal/localization"
)

func (h *Handler) handleUserStart(b *gotgbot.Bot, ctx *ext.Context) error {
	telegramID := ctx.Message.From.Id

	b.SendMessage(
		telegramID,
		fmt.Sprintf(localization.GetMessage("en", "startCMD.message"), ctx.Message.Text),
		nil,
	)
	return nil
}
