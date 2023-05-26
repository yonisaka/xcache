package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/yonisaka/cache/config"
	"github.com/yonisaka/cache/pkg/cache"
	"log"
	"time"
)

func main() {
	var (
		ctx         = context.Background()
		key         = "user:YSS2959"
		val         = "Hello Yonisaka"
		expiryCache = 120 * time.Second // 120 seconds
	)

	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.Initialize()

	redis, err := cache.NewRedisCache(conf)
	if err != nil {
		log.Fatal(err)
	}

	redisCacheInstance := cache.New(redis)

	errSet := redisCacheInstance.Set(ctx, key, val, expiryCache)
	if errSet != nil {
		log.Fatal(errSet)
	}

	valGet, errGet := redisCacheInstance.Get(ctx, key)
	if errGet != nil {
		log.Fatal(errGet)
	}

	log.Println("Get Redis Cache Successfully")
	log.Println(valGet)

	mc, err := cache.NewMemcachedCache(conf)
	if err != nil {
		log.Fatal(err)
	}

	mcCacheInstance := cache.New(mc)

	errSet = mcCacheInstance.Set(ctx, key, val, expiryCache)
	if errSet != nil {
		log.Fatal(errSet)
	}

	valGet, errGet = mcCacheInstance.Get(ctx, key)
	if errGet != nil {
		log.Fatal(errGet)
	}

	log.Println("Get Memcached Cache Successfully")
	log.Println(valGet)
}
