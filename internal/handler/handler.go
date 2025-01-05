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
	"github.com/morf1lo/tgbotbase/internal/config"
	"github.com/morf1lo/tgbotbase/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	ctx context.Context // For operations in handler's services
	logger   *zap.Logger
	services *service.Service // Your services from 'internal/service'
	admins   map[int64]bool // Telegram bot admin ids
	cfg      *config.BotConfig // Bot config from 'internal/config'
}

func New(ctx context.Context, logger *zap.Logger, services *service.Service, cfg *config.BotConfig) *Handler {
	i := 0
	isThereAnyAdmins := true
	admins := map[int64]bool{}
	for isThereAnyAdmins {
		adminID, err := strconv.Atoi(os.Getenv("ADMIN_ID_" + strconv.Itoa(i)))
		if err == nil {
			admins[int64(adminID)] = true
			i++
		} else {
			isThereAnyAdmins = false
		}
	}

	return &Handler{
		ctx: ctx,
		logger: logger,
		services: services,
		admins: admins,
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
	// ...

	// Callbacks
	// Or you can add a separate handler for each callback
	// Read documentation to github.com/PaulSonOfLars/gotgbot
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.All, h.CallbackQueryHandler))

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
