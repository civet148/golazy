package cmds

import (
	"github.com/civet148/golazy/utils"
	"github.com/urfave/cli/v2"
)

const (
	subCmd_RunDsh = "dsh"
)

var CmdRun = &cli.Command{
	Name:  "run",
	Usage: "run commands",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		cmdRunDsh,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

var cmdRunDsh = &cli.Command{
	Name:  subCmd_RunDsh,
	Usage: "run deepseek harness with npx",
	Flags: []cli.Flag{},
	Action: func(ctx *cli.Context) error {
		return utils.ShellExec("npx @deepseek-ai/dsh web")
	},
}
