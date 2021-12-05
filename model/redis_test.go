package model

import (
	"log"
	"short-url/config"
	"testing"

	"github.com/gomodule/redigo/redis"
)

const (
	key = iota
	value
)

var cases = []struct {
	input    [2]string
	expected [2]string
}{
	{[...]string{"key", "val"}, [...]string{"key", "val"}},
	{[...]string{"", ""}, [...]string{"", ""}},
}

func TestRedis(t *testing.T) {
	config := &config.Config
	// init Redis
	redisPool, err := InitRedis(config)
	if err != nil {
		log.Println("redis pool cannot be created")
	}
	conn := redisPool.Get()
	defer conn.Close()

	for _, c := range cases {
		_, err = conn.Do("SET", c.input[key], c.input[value])
		if err != nil {
			t.Fatalf("redis set fail %v", err)
		}
	}
	for _, c := range cases {
		val, err := redis.String(conn.Do("GET", c.expected[key]))
		if err != nil {
			t.Fatalf("redis get fail %v", err)
		}
		if val != c.expected[value] {
			t.Fatalf("Input %s. Expected: %s, actual: %s\n", c.input[value], c.expected[value], val)
		}
	}
}
