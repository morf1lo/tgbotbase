package redisrepo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Default interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	SetJSON(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
	Decr(ctx context.Context, key string) *redis.IntCmd
}

type Redis struct {
	Default
}

func New(rdb *redis.Client) *Redis {
	return &Redis{
		Default: newDefaultRepo(rdb),
	}
}
