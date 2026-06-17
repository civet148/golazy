package cmds

import (
	"github.com/civet148/golazy/gen"
	"github.com/civet148/golazy/parser"
	"github.com/urfave/cli/v2"
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

var CmdApi = &cli.Command{
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
			return err
		}
		return gen.GenerateGoCode(flags, services)
	},
}
