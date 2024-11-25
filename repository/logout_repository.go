package repository

import (
	"context"
	"os"

	"time"

	"github.com/redis/go-redis/v9"
)

type blacklistRepository struct {
	redis *redis.Client
}

// AddTokenToBlacklist implements BlacklistRepository.
func (b *blacklistRepository) AddTokenToBlacklist(token string) error {
	ctx := context.Background()

	expired := os.Getenv("TOKEN_EXPIRED")
	expiredTime, _ := time.ParseDuration(expired)

	return b.redis.Set(ctx, token, "blacklisted", expiredTime).Err()
}

// IsBlacklisted implements BlacklistRepository.
func (b *blacklistRepository) IsBlacklisted(token string) (bool, error) {
	ctx := context.Background()

	val, err := b.redis.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil // Token tidak ada di blacklist
	} else if err != nil {
		return false, err // Error di redis
	}

	return val == "blacklisted", nil
}

type BlacklistRepository interface {
	AddTokenToBlacklist(token string) error
	IsBlacklisted(token string) (bool, error)
}

func NewBlacklistRepository(redis *redis.Client) BlacklistRepository {
	return &blacklistRepository{redis: redis}
}
