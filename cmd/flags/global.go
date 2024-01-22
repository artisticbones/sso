package flags

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

// path is the SSO configuration file path
var (
	path      cli.Path
	logLevel  string
	jwtSecret string
	mode      string
)

var CfgFilePathFlag = cli.PathFlag{
	Name:        "config",
	DefaultText: ".\"/\"conf.yaml",
	Usage:       "Location of server config `FILE`",
	Value:       ".\"/\"conf.yaml",
	Destination: &path,
	EnvVars:     []string{"SSO_CFG"},
	Action:      nil,
}

var ModeFlag = cli.StringFlag{
	Name:        "mode",
	DefaultText: "prod",
	Usage:       "Switch server mode",
	Value:       "prod",
	Destination: &mode,
	EnvVars:     []string{"SSO_MODE"},
	Action: func(ctx *cli.Context, s string) error {

		return nil
	},
}

var LogLevelFlag = cli.StringFlag{
	Name:        "log-level",
	DefaultText: "info",
	Usage:       "Set the logging level (\"debug\", \"info\", \"warn\", \"error\", \"fatal\")",
	Value:       "info",
	Destination: &logLevel,
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

var JwtFlag = cli.StringFlag{}
