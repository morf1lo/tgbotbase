package handler

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/precheckoutquery"
	"github.com/morf1lo/tgbotbase/internal/config"
	"github.com/morf1lo/tgbotbase/internal/service"
	"go.uber.org/zap"
)

const (
	START_STATE = "start-state"
)

var c = context.Background()

type Handler struct {
	logger   *zap.Logger
	services *service.Service // Your services from 'internal/service'
	admins   map[int64]bool // Telegram bot admin ids
	cfg      *config.BotConfig // Bot config from 'internal/config'
}

func New(logger *zap.Logger, services *service.Service, cfg *config.BotConfig) *Handler {
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
		cfg: cfg,
	}
}

func (h *Handler) RunBot() {
	b, err := gotgbot.NewBot(h.cfg.Token, nil)
	if err != nil {
		h.logger.Sugar().Fatalf("failed to create bot: %s", err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			h.logger.Sugar().Errorf("an error occurred while handling update: %s", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	// Conversations
	dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCommand("start", h.StartCMD)},
		map[string][]ext.Handler{
			START_STATE: {handlers.NewMessage(h.noCommands, h.handleUserStart)},
		},
		&handlers.ConversationOpts{
			Exits: nil,
			StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySender),
			AllowReEntry: true,
		},
	))

	// Commands
	dispatcher.AddHandler(handlers.NewCommand("start", h.StartCMD))

	// Callbacks
	// Or you can add a separate handler for each callback
	// Read documentation to github.com/PaulSonOfLars/gotgbot/v2
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.All, h.CallbackQueryHandler))

	// Payments
	dispatcher.AddHandler(handlers.NewPreCheckoutQuery(precheckoutquery.All, h.PreCheckoutHandler))
	dispatcher.AddHandler(handlers.NewMessage(message.SuccessfulPayment, h.SuccessfulPaymentHandler))

	// Start polling
	if err := updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: false,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	}); err != nil {
		h.logger.Sugar().Fatalf("failed to start polling: %s", err.Error())
	}
}

// Filtering admins
func (h *Handler) IsAdmin(ctx *ext.Context) bool {
	return !ctx.Message.From.IsBot && h.admins[ctx.Message.From.Id]
}

func (h *Handler) noCommands(msg *gotgbot.Message) bool {
	return message.Text(msg) && !message.Command(msg)
}
