package handler

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (h *Handler) PreCheckoutHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	b.AnswerPreCheckoutQuery(ctx.PreCheckoutQuery.Id, true, nil)
	return nil
}
