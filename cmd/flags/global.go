package flags

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

// path is the SSO configuration file path
var (
	path      cli.Path
	logLevel  string
	logFile   string
	jwtSecret string
	mode      string
	orm       string
	cache     string
)

func Mode() string {
	return mode
}

func OrmUri() string {
	return orm
}

func CacheUri() string {
	return cache
}

func CfgPath() string {
	return path
}

func LogLevel() string {
	return logLevel
}

func LogFile() string {
	return logFile
}

func JwtSecret() string {
	return jwtSecret
}

var CfgFilePathFlag = cli.PathFlag{
	Name:        "config",
	Usage:       "Location of server config `FILE`",
	Destination: &path,
	Aliases:     []string{"c"},
	EnvVars:     []string{"SSO_CFG"},
}

var ModeFlag = cli.StringFlag{
	Name:        "mode",
	DefaultText: "prod",
	Usage:       "Switch server mode",
	Value:       "prod",
	Destination: &mode,
	Aliases:     []string{"m"},
	EnvVars:     []string{"SSO_MODE"},
	Action: func(ctx *cli.Context, s string) error {
		if s != "" {
			if s == "dev" || s == "prod" {
				return nil
			}
			return fmt.Errorf("mode %s is not support", s)
		}
		return nil
	},
}

var LogLevelFlag = cli.StringFlag{
	Name:        "log-level",
	DefaultText: "info",
	Usage:       "Set the logging level (\"debug\", \"info\", \"warn\", \"error\", \"fatal\")",
	Value:       "info",
	Destination: &logLevel,
	Aliases:     []string{"l"},
	EnvVars:     []string{"SSO_CFG"},
	Action: func(ctx *cli.Context, s string) error {
		levels := map[string]struct{}{
			"debug": struct{}{},
			"info":  struct{}{},
			"warn":  struct{}{},
			"error": struct{}{},
			"fatal": struct{}{},
		}
		if level := ctx.String("log-level"); level != "" {
			if _, ok := levels[level]; !ok {
				return fmt.Errorf("log-level %s is not support", level)
			}
		}
		return nil
	},
}

var JwtFlag = cli.StringFlag{
	Name:        "jwt-secret",
	DefaultText: "./jwt-secret",
	Usage:       "Jwt secret phrase for user login authorization",
	// Note that default values set from file (e.g. FilePath) take precedence over default values set from the environment (e.g. EnvVar).
	FilePath:    "./jwt-secret",
	Destination: &jwtSecret,
	EnvVars:     []string{"SSO_JWT_SECRET"},
	Action:      nil,
}

var OrmFlag = cli.StringFlag{
	Name:        "orm",
	DefaultText: "mysql://root:crane@127.0.0.1:3306/sso?charset=utf8mb4&parseTime=True&loc=UTC",
	Usage:       "`URI` of SSO background database",
	// Note that default values set from file (e.g. FilePath) take precedence over default values set from the environment (e.g. EnvVar).
	Value:       "mysql://root:crane@127.0.0.1:3306/sso?charset=utf8mb4&parseTime=True&loc=UTC",
	Destination: &orm,
	Action:      nil,
}

var CacheFlag = cli.StringFlag{
	Name:        "cache",
	DefaultText: "redis://127.0.0.1:6379/0?dial_timeout=5s&read_timeout=3s",
	Usage:       "`URI` of SSO background cache",
	// Note that default values set from file (e.g. FilePath) take precedence over default values set from the environment (e.g. EnvVar).
	Value:       "redis://127.0.0.1:6379/0?dial_timeout=5s&read_timeout=3s",
	Destination: &cache,
	Action:      nil,
}
