package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ruriazz/ocean-test/pkg/configs"
)

type RedisClient interface {
	GetString(key string) (string, error)
	SetString(key string, value string, expiration time.Duration) error
	Unset(key string) error
	Keys(prefix string) ([]string, error)
	TTL(key string) (time.Duration, error)
}

type redisClient struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedisClient(ctx context.Context, config configs.Config) (RedisClient, error) {
	c := &redisClient{
		ctx: ctx,
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
			Username: config.RedisUser,
			Password: config.RedisPass,
			DB:       0,
		}),
	}

	_, err := c.client.Ping(ctx).Result()
	return c, err
}

func (rc redisClient) GetString(key string) (string, error) {
	return rc.client.Get(rc.ctx, key).Result()
}

func (rc redisClient) SetString(key string, value string, expiration time.Duration) error {
	return rc.client.Set(rc.ctx, key, value, expiration).Err()
}

func (rc redisClient) Unset(key string) error {
	_, err := rc.client.Del(rc.ctx, key).Result()
	return err
}

func (rc redisClient) Keys(prefix string) ([]string, error) {
	var cursor uint64
	var keys []string

	if !strings.HasSuffix(prefix, "*") {
		prefix += "*"
	}

	for {
		res, newCursor, err := rc.client.Scan(rc.ctx, cursor, prefix, 1000).Result()
		if err != nil {
			return nil, err
		}

		keys = append(keys, res...)

		cursor = newCursor
		if cursor == 0 {
			break
		}
	}

	return keys, nil
}

func (rc redisClient) TTL(key string) (time.Duration, error) {
	ttl := rc.client.TTL(rc.ctx, key)
	if ttl.Err() != nil {
		return 0, ttl.Err()
	}

	return ttl.Val(), nil
}
