package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/tgbotbase/internal/model"
)

type Admin interface {
	Create(ctx context.Context, admin *model.Admin) error
	FindByTelegramID(ctx context.Context, telegramID int64) (*model.Admin, error)
}

type Postgres struct {
	Admin
}

func New(db *pgx.Conn) *Postgres {
	return &Postgres{
		Admin: newAdminRepo(db),
	}
}
