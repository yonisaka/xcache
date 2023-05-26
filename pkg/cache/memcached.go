package cache

import (
	"context"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/yonisaka/cache/config"
	"time"
)

// MemcachedCache is a type
type MemcachedCache struct {
	cfg    *config.Config
	client *memcache.Client
}

// Get is a method
func (m *MemcachedCache) Get(_ context.Context, key string) (interface{}, error) {
	item, err := m.client.Get(key)
	if err != nil {
		return nil, err
	}

	return string(item.Value), nil
}

// Set is a method
func (m *MemcachedCache) Set(_ context.Context, key string, val interface{}, exp time.Duration) error {
	if exp < 1 {
		exp = DefaultExpiration
	}

	expiration, err := time.ParseDuration(exp.String())
	if err != nil {
		return err
	}

	item := &memcache.Item{
		Key:        key,
		Value:      []byte(val.(string)),
		Expiration: int32(expiration.Seconds()),
	}

	return m.client.Set(item)
}

// Delete is a method
func (m *MemcachedCache) Delete(_ context.Context, key string) error {
	return m.client.Delete(key)
}

// NewMemcachedCache is a constructor
func NewMemcachedCache(cfg *config.Config) (*MemcachedCache, error) {
	dns := fmt.Sprintf("%s:%d", cfg.MemcachedConfig.Host, cfg.MemcachedConfig.Port)

	mc := memcache.New(dns)

	return &MemcachedCache{
		cfg:    cfg,
		client: mc,
	}, nil
}

var _ CacheStrategy = &MemcachedCache{}
