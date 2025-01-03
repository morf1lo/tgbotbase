package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/tgbotbase/internal/model"
)

type adminRepo struct {
	db *pgx.Conn
}

func newAdminRepo(db *pgx.Conn) Admin {
	return &adminRepo{
		db: db,
	}
}

func (r *adminRepo) Create(ctx context.Context, admin *model.Admin) error {
	_, err := r.db.Exec(ctx, "INSERT INTO admins(telegram_id) VALUES($1)", admin.TelegramID)
	return err
}

func (r *adminRepo) FindByTelegramID(ctx context.Context, telegramID int64) (*model.Admin, error) {
	var admin model.Admin
	if err := r.db.QueryRow(ctx, "SELECT a.id, a.telegram_id, a.created_at from admins a WHERE a.telegram_id = $1", telegramID).Scan(
		&admin.ID,
		&admin.TelegramID,
		&admin.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &admin, nil
}
