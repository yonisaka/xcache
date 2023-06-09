package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yonisaka/cache/config"
	"time"
)

// RedisCache is a type
type RedisCache struct {
	cfg    *config.Config
	client *redis.Client
}

// Get is a method
func (r *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	return r.client.Get(ctx, key).Result()
}

// Set is a method
func (r *RedisCache) Set(ctx context.Context, key string, val interface{}, exp time.Duration) error {
	return r.client.Set(ctx, key, val, exp).Err()
}

// Delete is a method
func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// NewRedisCache is a constructor
func NewRedisCache(cfg *config.Config) (*RedisCache, error) {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisConfig.Host, cfg.RedisConfig.Port),
		Password: cfg.RedisConfig.Password,
		DB:       cfg.RedisConfig.DB,
	}

	conn := redis.NewClient(opt)

	_, err := conn.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}

	return &RedisCache{
		cfg:    cfg,
		client: conn,
	}, nil
}

var _ CacheStrategy = &RedisCache{}
