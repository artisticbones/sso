package commands

import (
	"context"
	"fmt"
	"github.com/artisticbones/sso/configs"
	"github.com/artisticbones/sso/init/cache"
	"github.com/artisticbones/sso/init/database"
	"github.com/artisticbones/sso/init/gin"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

const (
	cliName        = "commands"
	cliUsage       = "[OPTIONS] COMMAND"
	cliDescription = "A simple commands line client for commands"
	copyright      = ``
)

var (
	path        = ""
	globalFlags = []cli.Flag{}
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
	db := database.GetDB(true)
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
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
	cfg := configs.Load(path)
	cache.Get(cfg.Cache.RedisUri())
	return nil
}

func init() {
	app.Commands = append(app.Commands)
	app.Flags = append(app.Flags, globalFlags...)
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
