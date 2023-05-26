package config

import (
	"os"
	"strconv"
)

type (
	Config struct {
		AppName         string
		AppPort         int
		RedisConfig     CacheConfig
		MemcachedConfig CacheConfig
	}

	CacheConfig struct {
		Host     string
		Port     int
		Password string
		DB       int
	}
)

func Initialize() *Config {
	return &Config{
		AppName: getEnv("APP_NAME", "go-boilerplate"),
		AppPort: getEnvAsInt("APP_PORT", 8080),
		RedisConfig: CacheConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		MemcachedConfig: CacheConfig{
			Host: getEnv("MEMCACHED_HOST", "localhost"),
			Port: getEnvAsInt("MEMCACHED_PORT", 11211),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultVal
}
