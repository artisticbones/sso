package apis

import (
	"context"
	"github.com/artisticbones/sso/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

type API struct {
	Context *gin.Context
	Logger  *logger.Logger
	Cache   *redis.Client
}

const (
	defaultCacheTime = 15 * time.Minute
)

var prefix = func() string { return "" }

func (api API) Set(key string, val interface{}) error {
	return api.Cache.Set(context.Background(), prefix()+key, val, defaultCacheTime).Err()
}
