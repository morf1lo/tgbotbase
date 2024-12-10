package redisrepo

import "github.com/redis/go-redis/v9"

type Redis struct {
}

func New(rdb *redis.Client) *Redis {
	return &Redis{
	}
}
