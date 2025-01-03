package service

import (
	"context"

	"github.com/morf1lo/tgbotbase/internal/model"
	"github.com/morf1lo/tgbotbase/internal/repository"
)

type Admin interface {
	Create(ctx context.Context, admin *model.Admin) error
	FindByTelegramID(ctx context.Context, telegramID int64) (*model.Admin, error)
}

type Service struct {
	Admin
}

func New(repo *repository.Repository) *Service {
	return &Service{
		
	}
}
