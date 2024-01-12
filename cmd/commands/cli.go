package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

const (
	cliName        = "commands"
	cliUsage       = "[OPTIONS] COMMAND"
	cliDescription = "A simple commands line client for commands"

	defaultDialTimeout      = 2 * time.Second
	defaultCommandTimeOut   = 5 * time.Second
	defaultKeepAliveTime    = 2 * time.Second
	defaultKeepAliveTimeOut = 6 * time.Second
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

var (
	app = &cli.App{
		Name:                   cliName,
		Usage:                  cliUsage,
		Args:                   true,
		Description:            cliDescription,
		EnableBashCompletion:   true,
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		CommandNotFound:        commandNotFound,
		Compiled:               compileTime(),
		Authors:                authors,
		Copyright:              "",
		Reader:                 os.Stdin,
		Writer:                 os.Stdout,
		ErrWriter:              os.Stderr,
		UseShortOptionHandling: true,
		Suggest:                true,
	}
)

func init() {
	app.Commands = append(app.Commands)
}

func Start() error {
	// Make help just show the usage
	return app.Run(os.Args)
}

func MustStart() {
	if err := Start(); err != nil {
		log.Fatalln(err)
	}
}
