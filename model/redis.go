package model

import (
	"context"
	"short-url/config"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type RedisPool struct {
	pool *redis.Pool
}

// InitRedis for reids initialization
func InitRedis(config *config.Configure) (*RedisPool, error) {
	ctx := context.Background()
	pool, err := newRdbPool(ctx, config)
	if err != nil {
		logrus.Error(err)
	}
	return &RedisPool{
		pool: pool,
	}, nil
}

func newRdbPool(ctx context.Context, config *config.Configure) (*redis.Pool, error) {
	pool := redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", *config.Redis.Address)
			if err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return &pool, nil
}

func (r *RedisPool) Close() error {
	return r.pool.Close()
}

func (r *RedisPool) Get() redis.Conn {
	return r.pool.Get()
}
