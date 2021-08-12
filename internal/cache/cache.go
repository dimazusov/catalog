package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"catalog/internal/pkg/apperror"
)

const lifetime = 8 * time.Hour

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (value interface{}, err error)
}

type cache struct {
	redis *redis.Client
}

func New(redis *redis.Client) Cache {
	return &cache{redis: redis}
}

func (m cache) Set(key string, value interface{}) error {
	err := m.redis.Set(context.Background(), key, value, lifetime).Err()
	if err != nil {
		return errors.Wrapf(err, "cannot set cache %s %#v", key, value)
	}
	return nil
}

func (m cache) Get(key string) (value interface{}, err error) {
	value, err = m.redis.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrapf(err, "cannot get cache %s", key)
	}

	return value, nil
}
