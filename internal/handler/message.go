package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) HandleMessage(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Handling message using 'internal/model/State'
	// statesFunctions := map[model.State]func(tgbot *tgbotapi.BotAPI, update tgbotapi.Update){
    //     model.NilState: h.handleNilState,
    //     model.WaitingForNextMessageState: h.handleNextMessage,
    // }

    // statesFunctions[h.states.CurrentState](tgbot, update)
}

func (h *Handler) handleNilState(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
}

func (h *Handler) handleNextMessage(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
}
