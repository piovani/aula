package cache

import (
	"context"
	"time"

	"github.com/piovani/aula/infra/config"
	redis "github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.ClusterClient
}

func NewCache() *Cache {
	return &Cache{
		client: redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{config.Env.RedisURL},
		}),
	}
}

func (c *Cache) Set(key string, value string, ttl time.Duration) error {
	return c.client.Set(context.Background(), key, value, ttl).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *Cache) Delete(key string) error {
	return c.client.Del(context.Background(), key).Err()
}
