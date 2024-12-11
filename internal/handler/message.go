package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) HandleMessage(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Handle message using current user state from Redis

    // userID := update.Message.From.ID

    // if (h.HasPermission(update)) {
    //     adminState, err := h.services.Cache.Get(ctx, fmt.Sprintf("%d:state", userID)).Result()
    //     if err != nil && err != redis.Nil {
    //         h.logger.Sugar().Errorf("failed to get admin(%d) state from redis: %s", userID, err.Error())
    //         return
    //     }
    //     if err == redis.Nil {
    //         adminState = model.NilState
    //     }
    
    //     adminStatesFunctions := map[string]func(tgbot *tgbotapi.BotAPI, update tgbotapi.Update){
    //         model.NilState: h.handleNilState,
    //         model.WaitingForNextMessageState: h.handleNextMessage,
    //     }

    //     h.logger.Sugar().Infof("Admin(%d) current state is %s", userID, adminState)
    //     adminStatesFunctions[adminState](tgbot, update)
    // } else {
            // Handle NOT admin message
    // }
}

func (h *Handler) handleNilState(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
}

func (h *Handler) handleNextMessage(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
}
