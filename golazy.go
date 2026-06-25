package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/civet148/golazy/cmds"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
)

var (
	Version     = "v0.9.3"
	ProgramName = "golazy"
	BuildTime   = "2026-06-25"
	GitCommit   = "<N/A>"
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

func main() {
	grace()
	app := &cli.App{
		Name:    ProgramName,
		Usage:   fmt.Sprintf("%s <sub-command> [options] ", ProgramName),
		Version: fmt.Sprintf("%s %s commit %s", Version, BuildTime, GitCommit),
		Commands: []*cli.Command{
			cmds.CmdApi,
			cmds.CmdInstall,
			cmds.CmdGen,
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
