package service

import (
	"context"
	"time"

	"github.com/morf1lo/tgbotbase/internal/model"
	"github.com/morf1lo/tgbotbase/internal/repository"
	"github.com/morf1lo/tgbotbase/internal/repository/redisrepo"
	"github.com/redis/go-redis/v9"
)

type adminService struct {
	repo *repository.Repository
}

func newAdminService(repo *repository.Repository) Admin {
	return &adminService{
		repo: repo,
	}
}

func (s *adminService) Create(ctx context.Context, admin *model.Admin) error {
	return s.repo.Postgres.Admin.Create(ctx, admin)
}

func (s *adminService) FindByTelegramID(ctx context.Context, telegramID int64) (*model.Admin, error) {
	adminCache, err := redisrepo.Get[model.Admin](s.repo.Default, ctx, redisrepo.AdminKey(telegramID))
	if err == nil {
		return adminCache, nil
	}
	if err != redis.Nil {
		return nil, err
	}

	admin, err := s.repo.Postgres.Admin.FindByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Redis.Default.SetJSON(ctx, redisrepo.AdminKey(telegramID), admin, time.Hour); err != nil {
		return nil, err
	}

	return admin, nil
}
