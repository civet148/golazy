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
	subCmd_DB2GO       = "db2go"
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
		subCmdDB2GO,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

func generateFile(outputDir, outputName string, data []byte) error {
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
	outputName = filepath.Join(outputDir, outputName)
	if err = os.WriteFile(outputName, data, 0755); err != nil {
		return fmt.Errorf("❌ 写入文件%s错误: %s", outputName, err)
	}
	log.Printf("✅ 创建文件%s成功", outputName)
	return nil
}

//go:embed tpls/genproto.tpl
var genProtoTemplate string

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
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(genProtoTemplate))
	},
}

//go:embed tpls/db2go.tpl
var db2goTemplate string

var subCmdDB2GO = &cli.Command{
	Name:  subCmd_DB2GO,
	Usage: "generate db2go script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "db2go script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "db2go script output file name",
			Value:   "db2go.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(db2goTemplate))
	},
}
