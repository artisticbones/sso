package configs

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const (
	defaultDialTimeout      = 10 * time.Second
	defaultCommandTimeOut   = 5 * time.Second
	defaultKeepAliveTime    = 2 * time.Second
	defaultKeepAliveTimeOut = 6 * time.Second
)

type Backend interface {
	Uri() string
}

type RedisConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

// Uri redis :// [[username :] password@] host [:port][/database]
//
//	[?[timeout=timeout[d|h|m|s|ms|us|ns]] [&clientName=clientName]
//	[&libraryName=libraryName] [&libraryVersion=libraryVersion] ]
func (cfg *RedisConfig) Uri() string {
	uri := "redis://"
	if cfg.User != "" {
		uri = uri + cfg.User + ":"
	}
	if cfg.Password != "" {
		uri = uri + cfg.Password + "@"
	}
	return fmt.Sprintf("%s?dial_timeout=%ds&read_timeout=%ds", uri+cfg.Host+":"+cfg.Port+"/"+cfg.Database, defaultDialTimeout/time.Second, defaultCommandTimeOut/time.Second)
}

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Net      string `yaml:"net"`
}

func (cfg *MysqlConfig) Uri() string {
	uri := ""
	if cfg.User != "" {
		uri = uri + cfg.User + ":"
	}
	if cfg.Password != "" {
		uri = uri + cfg.Password + "@"
	}
	return uri + cfg.Net + "(" + cfg.Host + ":" + cfg.Port + ")" + "/" + cfg.Database + "?charset=utf8mb4&parseTime=True&loc=UTC"
}

type logger struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type Config struct {
	Mode      string      `yaml:"mode"`
	JwtSecret string      `yaml:"jwtSecret"`
	Log       logger      `yaml:"log"`
	Orm       MysqlConfig `yaml:"orm"`
	Cache     RedisConfig `yaml:"cache"`
}

var (
	_cfg *Config
	once sync.Once
)

func New(mode, jwt, level, file, orm, cache string) *Config {
	config, err := mysql.ParseDSN(orm)
	if err != nil {
		return nil
	}
	options, err := redis.ParseURL(cache)
	if err != nil {
		return nil
	}

	once.Do(func() {
		_cfg = &Config{
			Mode:      mode,
			JwtSecret: jwt,
			Log: logger{
				Level: level,
				File:  file,
			},
			Orm: MysqlConfig{
				User:     config.User,
				Password: config.Passwd,
				Host:     config.Addr,
				Port:     "3306",
				Database: config.DBName,
				Net:      config.Net,
			},
			Cache: RedisConfig{
				User:     options.Username,
				Password: options.Password,
				Host:     options.Addr,
				Port:     "6379",
				Database: fmt.Sprintf("%d", options.DB),
			},
		}
	})

	return _cfg
}

func Load(path string) *Config {
	once.Do(func() {
		var cfg = &Config{}
		file, err := os.Open(path)
		if err != nil {
			log.Print(err)
			return
		}
		body, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = yaml.Unmarshal(body, cfg)
		if err != nil {
			log.Fatal(err)
			return
		}
		_cfg = cfg
	})
	return _cfg
}

func Global() *Config {
	return _cfg
}
