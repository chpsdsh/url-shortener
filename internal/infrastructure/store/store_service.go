package store

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

const CacheDuration = 6 * time.Hour

var ctx = context.Background()

type StorageService struct {
	redisClient *redis.Client
}

func NewStorageService() (StorageService, error) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		return StorageService{}, fmt.Errorf("error init Redis: %v", err)
	}

	fmt.Println("redis started successfully", pong)
	return StorageService{redisClient: redisClient}, nil
}

func (service *StorageService) Close() {
	service.redisClient.Close()
}

func (s StorageService) SaveUrlMapping(shortUrl, originalUrl string) error {
	if err := s.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err(); err != nil {
		return fmt.Errorf("error saving url-shortener: %v", err)
	}
	return nil
}

func (s StorageService) RetrieveOriginalUrl(shortUrl string) (string, error) {
	res, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", fmt.Errorf("error getting originalUrl: %v", err)
	}
	return res, nil
}
