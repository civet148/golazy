package cmds

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
)

const (
	Cmd_Gen = "gen"
)

const (
	subCmd_ProtoScript = "proto-script"
)

const (
	cmdFlag_Output = "output"
	cmdFlag_Name   = "name"
)

var CmdGen = &cli.Command{
	Name:  Cmd_Gen,
	Usage: "generation commands",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		subCmdProtoScript,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

var subCmdProtoScript = &cli.Command{
	Name:  subCmd_ProtoScript,
	Usage: "generate protobuf compile script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "protobuf generation script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "protobuf generation script output file name",
			Value:   "genproto",
		},
	},
	Action: func(ctx *cli.Context) error {
		outputDir := ctx.String(cmdFlag_Output)
		if outputDir == "" {
			// 判断$GOPATH是否存在，如果存在则使用$GOPATH/src目录作为下载基础目录
			gopath := os.Getenv("GOPATH")
			if gopath == "" {
				return fmt.Errorf("GOPATH environment variable is not set")
			}
			outputDir = filepath.Join(gopath, "bin")
		}
		var err error
		if err = os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return fmt.Errorf("❌ 创建目录%s失败: %s", outputDir, err)
		}
		outputName := filepath.Join(outputDir, ctx.String(cmdFlag_Name))
		if err = os.WriteFile(outputName, []byte(genProtoTemplate), 0755); err != nil {
			return fmt.Errorf("write file %s error: %s", outputName, err)
		}
		log.Printf("✅ 创建文件%s成功", outputName)
		return nil
	},
}

//go:embed tpls/genproto.tpl
var genProtoTemplate string
