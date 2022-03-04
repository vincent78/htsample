package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"htSample/cmd"
	"htSample/global"
	"os"
	"path/filepath"
	"sort"
)

var (
	DefaultFlags = []cli.Flag{
		DBHostFlag,
		DBPortFlag,
		DBNameFlag,
		DBUserFlag,
		DBPwdFlag,
	}

	//LogHomeFlag = &cli.StringFlag{
	//	Name:    "logHome",
	//	Aliases: []string{"l"},
	//	EnvVars: []string{global.EnvName("LOG_HOME")},
	//	Value:   "/opt/phoenix-qsh/logs",
	//	Usage:   "the path for log files ",
	//}
	//
	//LogFileFlag = &cli.StringFlag{
	//	Name:    "logFile",
	//	EnvVars: []string{global.EnvName("LOG_FILE")},
	//	Value:   "qsh",
	//	Usage:   "the log file name ",
	//}
	//
	//LogLevelFlag = &cli.IntFlag{
	//	Name:    "logLevel",
	//	EnvVars: []string{global.EnvName("LOG_LEVEL")},
	//	Value:   0,
	//	Usage:   "0 DEBUG,1 INFO,2 WARN,3 ERROR",
	//}
	//
	//DebugFlag = &cli.BoolFlag{
	//	Name:    "debug",
	//	EnvVars: []string{global.EnvName("DEBUG")},
	//	Value:   false,
	//	Usage:   "是否调试模式",
	//}

	DBHostFlag = &cli.StringFlag{
		Name:    "dbHost",
		EnvVars: []string{global.EnvName("DB_HOST")},
		Value:   "dev.vincent78.top",
		Usage:   "host of the db",
	}

	DBPortFlag = &cli.IntFlag{
		Name:    "dbPort",
		EnvVars: []string{global.EnvName("DB_PORT")},
		Value:   15432,
		Usage:   "port of the db",
	}

	DBNameFlag = &cli.StringFlag{
		Name:    "dbName",
		EnvVars: []string{global.EnvName("DB_NAME")},
		Value:   "htsample",
		Usage:   "name of the db",
	}

	DBUserFlag = &cli.StringFlag{
		Name:    "dbUser",
		EnvVars: []string{global.EnvName("DB_USER")},
		Value:   "htsample",
		Usage:   "user of the db",
	}

	DBPwdFlag = &cli.StringFlag{
		Name:    "dbPwd",
		EnvVars: []string{global.EnvName("DB_PWD")},
		Value:   "htsample",
		Usage:   "pwd of the db",
	}
)

func init() {
}

func main() {
	app := NewApp(global.GitCommit, global.BuildDate)
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func NewApp(gitCommit, buildDate string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = cmd.VersionWithCommit(gitCommit, buildDate)
	app.Action = cmd.Default
	app.HideVersion = true
	app.Usage = ""
	app.Copyright = "Copyright 2021-2026 The dataTube Authors"

	app.Commands = []*cli.Command{
		// 所有的Command必须在这里进行注册
		cmd.VersionCommand,
		cmd.ServerCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Flags = append(app.Flags, DefaultFlags...)
	sort.Sort(cli.FlagsByName(app.Flags))

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
	}

	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		if isSubcommand {
			return err
		}
		fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
		return nil
	}

	app.Before = func(ctx *cli.Context) error {
		return nil
	}
	app.After = func(ctx *cli.Context) error {
		return nil
	}

	return app
}
