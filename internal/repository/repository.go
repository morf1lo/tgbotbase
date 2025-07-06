package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/tgbotbase/internal/repository/postgres"
	"github.com/morf1lo/tgbotbase/internal/repository/redisrepo"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Postgres *postgres.Postgres
	Redis *redisrepo.Redis
}

func New(db *pgx.Conn, rdb *redis.Client) *Repository {
	return &Repository{
		Postgres: postgres.New(db),
		Redis: redisrepo.New(rdb),
	}
}
