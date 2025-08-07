package main

import (
	"fmt"
	"github.com/civet148/golazy/gen"
	"github.com/civet148/golazy/parser"
	"github.com/civet148/golazy/version"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

var (
	BuildTime = "2025-08-07"
	GitCommit = "<N/A>"
)

const (
	Cmd_Api = "api"
)

const (
	SubCmd_Go = "go"
)

const (
	CmdFlag_ApiFile = "api"
	CmdFlag_Dir     = "dir"
	CmdFlag_Style   = "style"
)

func init() {
	log.SetLevel("info")
}

func grace() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	go func() {
		for {
			select {
			case s := <-sigChannel:
				{
					if s != nil && s == os.Interrupt {
						fmt.Printf("Ctrl+C signal captured, program exiting...\n")
						close(sigChannel)
						os.Exit(0)
					}
				}
			}
		}
	}()
}

var cmdApi = &cli.Command{
	Name:  Cmd_Api,
	Usage: "api commands",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		subCmdGo,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

var subCmdGo = &cli.Command{
	Name:  SubCmd_Go,
	Usage: "generate go code from api file",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_ApiFile,
			Aliases: []string{"f"},
			Usage:   "api file path",
			Value:   "test.api",
		},
		&cli.StringFlag{
			Name:    CmdFlag_Dir,
			Aliases: []string{"o"},
			Usage:   "output directory",
			Value:   ".",
		},
		&cli.StringFlag{
			Name:    CmdFlag_Style,
			Aliases: []string{"s"},
			Usage:   "code style (go_lazy/goLazy/GoLazy)",
			Value:   "go_lazy",
		},
	},
	Action: func(ctx *cli.Context) error {
		flags := &gen.Config{
			ApiFile: ctx.String(CmdFlag_ApiFile),
			OutDir:  ctx.String(CmdFlag_Dir),
			Style:   ctx.String(CmdFlag_Style),
		}

		// 从文件加载并解析api内容
		services, err := parser.ParseApiFile(ctx.String(CmdFlag_ApiFile))
		if err != nil {
			log.Errorf(err.Error())
			return err
		}
		return gen.GenerateGoCode(flags, services)
	},
}

func main() {
	grace()
	app := &cli.App{
		Name:    version.ProgramName,
		Usage:   fmt.Sprintf("%s <sub-command> [options] ", version.ProgramName),
		Version: fmt.Sprintf("v%s %s commit %s", version.Version, BuildTime, GitCommit),
		Commands: []*cli.Command{
			cmdApi,
		},
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}
