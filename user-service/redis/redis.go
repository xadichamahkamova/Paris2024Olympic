package redis

import (
	"context"
	"fmt"
	config "user-service/internal/user/pkg/load"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg config.Config) (*redis.Client, error) {
	target := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: target,
	})

	if err := rdb.Ping(context.Background()); err != nil {
		return rdb, nil
	}

	return rdb, nil

}
