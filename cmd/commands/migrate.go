package commands

import (
	"fmt"
	"github.com/artisticbones/sso/cmd/flags"
	"github.com/artisticbones/sso/init/database"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var sqlPath cli.Path

var Migrate = cli.Command{
	Name:        "migrate",
	Usage:       "migrate tables to database you select",
	UsageText:   "",
	Description: "",
	Category:    "Database Commands",
	Action:      action,
	Flags:       []cli.Flag{&flags.OrmFlag, &sqlFlag, &autoMigrate},
}

var autoMigrate = cli.BoolFlag{
	Name:               "auto",
	Usage:              "enable auto migrate",
	DisableDefaultText: true,
}

var sqlFlag = cli.PathFlag{
	Name:        "sql",
	Usage:       "`Location of sql `FILE`",
	Destination: &sqlPath,
	Action: func(ctx *cli.Context, path cli.Path) error {
		if path == "" {
			return nil
		}
		_, err := os.Stat(sqlPath)
		return err
	},
}

var action = func(ctx *cli.Context) error {
	// init database
	db := database.GetDB(flags.OrmUri(), true)
	if ctx.Bool("auto") {
		fmt.Println("enable auto migrate!")
		err := db.AutoMigrate()
		if err != nil {
			return err
		}
	}
	// read file and execute
	if sqlPath != "" {
		body, err := os.ReadFile(sqlPath)
		if err != nil {
			return err
		}
		sqls := strings.Split(string(body), ";")
		for _, sql := range sqls {
			trimSql := strings.TrimSpace(sql)
			fmt.Println("Execute sql: ", trimSql)
			err := db.DB.Exec(trimSql)
			if err != nil {
				fmt.Println("Execute sql error: ", err.Error)
			}
		}
	}

	return db.Close()
}
