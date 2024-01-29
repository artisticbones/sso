package apis

import (
	"context"
	"github.com/artisticbones/sso/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type API struct {
	Context *context.Context
	Logger  *logger.Logger
	Cache   *redis.Client
}
