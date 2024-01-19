package configs

import (
	"fmt"
	"sync"
	"time"
)

const (
	defaultDialTimeout      = 10 * time.Second
	defaultCommandTimeOut   = 5 * time.Second
	defaultKeepAliveTime    = 2 * time.Second
	defaultKeepAliveTimeOut = 6 * time.Second
)

type UriConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

// RedisUri redis :// [[username :] password@] host [:port][/database]
//
//	[?[timeout=timeout[d|h|m|s|ms|us|ns]] [&clientName=clientName]
//	[&libraryName=libraryName] [&libraryVersion=libraryVersion] ]
func (cfg *UriConfig) RedisUri() string {
	uri := "redis://"
	if cfg.User != "" {
		uri = uri + cfg.User + ":"
	}
	if cfg.Password != "" {
		uri = uri + cfg.Password + "@"
	}
	return fmt.Sprintf("%s?dial_timeout=%ds&read_timeout=%ds", uri+cfg.Host+":"+cfg.Port+"/"+cfg.Database, defaultDialTimeout, defaultCommandTimeOut)
}

func (cfg *UriConfig) MysqlUri() string {
	uri := "mysql://"
	if cfg.User != "" {
		uri = uri + cfg.User + ":"
	}
	if cfg.Password != "" {
		uri = uri + cfg.Password + "@"
	}
	return uri + cfg.Host + ":" + cfg.Port + "/" + cfg.Database
}

type Config struct {
	Mode      string    `yaml:"mode"`
	JwtSecret string    `yaml:"jwtSecret"`
	LogLevel  string    `yaml:"logLevel"`
	Orm       UriConfig `yaml:"orm"`
	Cache     UriConfig `yaml:"cache"`
	once      sync.Once
}

func Load(path string) *Config {
	var cfg = &Config{}
	cfg.once.Do(func() {

	})
	return cfg
}
