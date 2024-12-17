package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) HandleMessage(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Handle message using current user state from Redis

	// userID := update.Message.From.ID

	// if h.IsAdmin(update) {
	// 	adminState, err := h.getUserState(userID)
	// 	if err != nil && err != redis.Nil {
	// 		h.logger.Sugar().Errorf("failed to get admin(%d) state from redis: %s", userID, err.Error())
	// 		return
	// 	}

	// 	adminStatesFunctions := map[string]func(tgbot *tgbotapi.BotAPI, update tgbotapi.Update){
	// 		model.NilState: h.adminHandleNilState,
	// 	}

	// 	adminStatesFunctions[adminState](tgbot, update)
	// } else {
	// 	userState, err := h.getUserState(userID)
	// 	if err != nil {
	// 		h.logger.Sugar().Errorf("failed to get user(%d) state from redis: %s", userID, err.Error())
	// 		return
	// 	}

	// 	userStatesFunctions := map[string]func(tgbot *tgbotapi.BotAPI, update tgbotapi.Update){
	// 		model.NilState: h.userHandleNilState,
	// 	}

	// 	userStatesFunctions[userState](tgbot, update)
	// }
}

// func (h *Handler) getUserState(userID int64) (string, error) {
// 	userState, err := h.services.Cache.Get(ctx, redisrepo.UserStateKey(userID)).Result()
// 	if err != nil && err != redis.Nil {
// 		return "", err
// 	}
// 	if err == redis.Nil {
// 		userState = model.NilState
// 	}

// 	return userState, nil
// }
