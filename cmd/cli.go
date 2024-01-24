package main

import (
	"context"
	"fmt"
	"github.com/artisticbones/sso/cmd/flags"
	"github.com/artisticbones/sso/configs"
	"github.com/artisticbones/sso/init/cache"
	"github.com/artisticbones/sso/init/database"
	"github.com/artisticbones/sso/init/gin"
	"github.com/artisticbones/sso/server/models/address"
	"github.com/artisticbones/sso/server/models/global"
	"github.com/artisticbones/sso/server/models/profile"
	"github.com/artisticbones/sso/server/models/user"
	"github.com/redis/go-redis/v9"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
	"time"
)

const (
	cliName        = "sso"
	cliUsage       = "[OPTIONS] COMMAND"
	cliDescription = "sso commands line client for commands"
	copyright      = ``
)

var authors = []*cli.Author{
	{Name: "Crane", Email: "artisticbones@163.com"},
}

// Tries to find out when this binary was compiled.
// Returns the current time if it fails to find it.
func compileTime() time.Time {
	info, err := os.Stat(os.Args[0])
	if err != nil {
		return time.Now()
	}
	return info.ModTime()
}

func commandNotFound(cCtx *cli.Context, command string) {
	fmt.Printf("sso: '%s' is not a sso command.\n", command)
	fmt.Printf("See 'sso --help'")
}

func action(cCtx *cli.Context) error {
	var (
		globalCfg = configs.Global()
	)
	f := func() (string, bool) {
		uri := globalCfg.Orm.Uri()
		if globalCfg.Mode == "dev" {
			return uri, true
		}
		return uri, false
	}
	clo := func(db *database.DB) {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}

	db := database.GetDB(f())
	defer clo(db)
	err := db.AutoMigrate(&address.Address{}, &global.AuthLog{}, &global.OptLog{}, &profile.Profile{}, &user.User{})
	if err != nil {
		return err
	}

	return gin.Start(cCtx.Context)
}

var (
	app = &cli.App{
		Name:                   cliName,
		Usage:                  cliUsage,
		Args:                   true,
		Description:            cliDescription,
		EnableBashCompletion:   true,
		Before:                 before,
		After:                  nil,
		Action:                 action,
		CommandNotFound:        commandNotFound,
		Compiled:               compileTime(),
		Authors:                authors,
		Copyright:              copyright,
		Reader:                 os.Stdin,
		Writer:                 os.Stdout,
		ErrWriter:              os.Stderr,
		UseShortOptionHandling: true,
		Suggest:                true,
	}
)

func before(ctx *cli.Context) error {
	// init config
	var (
		cfg *configs.Config
		rdb *redis.Client
	)

	if path := flags.CfgPath(); path != "" {
		cfg = configs.Load(path)
	} else {
		cfg = configs.New(flags.Mode(), flags.JwtSecret(), flags.LogLevel(), flags.LogFile(), flags.OrmUri(), flags.CacheUri())
	}
	if cfg == nil {
		return fmt.Errorf("cannot get config file")
	}

	rdb = cache.Get(cfg.Cache.Uri())

	resp, err := rdb.Ping(ctx.Context).Result()
	if err != nil {
		return err
	}
	if resp != "PONG" {
		return fmt.Errorf("ping rdb server but without %s response", resp)
	}
	return nil
}

func init() {
	app.Commands = append(app.Commands)
	// global flags
	app.Flags = append(app.Flags, &flags.ModeFlag, &flags.JwtFlag, &flags.LogLevelFlag, &flags.CfgFilePathFlag, &flags.OrmFlag, &flags.CacheFlag)
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))
}

func Start() error {
	// Make help just show the usage
	go cache.KeepAlive(context.Background())
	return app.Run(os.Args)
}

func MustStart() {
	if err := Start(); err != nil {
		log.Fatalln(err)
	}
}
