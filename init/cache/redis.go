package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

var (
	_rdb *redis.Client
	once sync.Once
	mu   sync.Mutex
)

func newCache() (*redis.Client, error) {
	options, err := redis.ParseURL("")
	if err != nil {
		return nil, err
	}
	return redis.NewClient(options), nil
}

func Get() *redis.Client {
	f := func() {
		rdb, err := newCache()
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
