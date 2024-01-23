package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

var (
	_rdb *redis.Client
	once sync.Once
	mu   sync.Mutex
)

func newCache(uri string) (*redis.Client, error) {
	options, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(options)
	resp, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	if resp != "PONG" {
		return nil, fmt.Errorf("cannot get %s from redis", resp)
	}
	return client, nil
}

func Get(uri string) *redis.Client {
	f := func() {
		rdb, err := newCache(uri)
		if err != nil {
			log.Fatalf("init redis error, err = %v", err)
		}
		_rdb = rdb
	}
	once.Do(f)

	mu.Lock()
	defer mu.Unlock()
	return _rdb
}

func Ping(ctx context.Context) string {
	return _rdb.Ping(ctx).Val()
}

func KeepAlive(ctx context.Context) {
	var (
		tick = time.NewTicker(60 * time.Second)
	)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case t := <-tick.C:
			if Ping(ctx) == "PONG" {
				continue
			}
			// need to alert
			fmt.Printf("cannot get response from redis. Time: %v", t)
		}
	}
}
