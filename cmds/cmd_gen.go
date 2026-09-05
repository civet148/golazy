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
	cmdFlag_Output = "output"
	cmdFlag_Name   = "name"
)

var CmdGen = &cli.Command{
	Name:  Cmd_Gen,
	Usage: "generation commands",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		cmdGenProto,
		cmdGenDB2GO,
		cmdGenMysql,
		cmdGenRedis,
		cmdGenRabbitmq,
		cmdGenMinio,
		cmdGenPostgres,
		cmdGenKafka,
		cmdGenInfluxdb,
		cmdGenProtoc,
		cmdGenDocker,
		cmdGenGit,
		cmdGenGolang,
		cmdGenDsh,
		cmdGenNvm,
		cmdGenK3S,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

func generateFile(outputDir, outputName string, data []byte) error {
	if outputDir == "" {
		// 判断$GOPATH是否存在，如果存在则使用$GOPATH/src目录作为下载基础目录
		gopath := os.Getenv("GOPATH")
		if gopath != "" {
			outputDir = filepath.Join(gopath, "bin")
		} else {
			outputDir = "."
		}
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

//go:embed tpls/docker.tpl
var genDockerTemplate string

var cmdGenDocker = &cli.Command{
	Name:  "docker",
	Usage: "generate docker install script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "docker.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(genDockerTemplate))
	},
}

//go:embed tpls/genproto.tpl
var genProtoTemplate string

var cmdGenProto = &cli.Command{
	Name:  "protobuf",
	Usage: "generate protobuf compile script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "genproto",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(genProtoTemplate))
	},
}

//go:embed tpls/db2go.tpl
var db2goTemplate string

var cmdGenDB2GO = &cli.Command{
	Name:  "db2go",
	Usage: "generate db2go script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "db2go.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(db2goTemplate))
	},
}

//go:embed tpls/mysql.tpl
var mysqlTemplate string

var cmdGenMysql = &cli.Command{
	Name:  "mysql",
	Usage: "generate mysql script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "mysql.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(mysqlTemplate))
	},
}

//go:embed tpls/redis.tpl
var redisTemplate string

var cmdGenRedis = &cli.Command{
	Name:  "redis",
	Usage: "generate redis script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "redis.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(redisTemplate))
	},
}

//go:embed tpls/rabbitmq.tpl
var rabbitmqTemplate string

var cmdGenRabbitmq = &cli.Command{
	Name:  "rabbitmq",
	Usage: "generate rabbitmq script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "rabbitmq.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(rabbitmqTemplate))
	},
}

//go:embed tpls/minio.tpl
var minioTemplate string

var cmdGenMinio = &cli.Command{
	Name:  "minio",
	Usage: "generate minio script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "minio.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(minioTemplate))
	},
}

//go:embed tpls/postgres.tpl
var postgresTemplate string

var cmdGenPostgres = &cli.Command{
	Name:  "postgres",
	Usage: "generate postgres script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "postgres.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(postgresTemplate))
	},
}

//go:embed tpls/kafka.tpl
var kafkaTemplate string

var cmdGenKafka = &cli.Command{
	Name:  "kafka",
	Usage: "generate kafka script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "kafka.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(kafkaTemplate))
	},
}

//go:embed tpls/influxdb.tpl
var influxdbTemplate string

var cmdGenInfluxdb = &cli.Command{
	Name:  "influxdb",
	Usage: "generate influxdb script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "influxdb.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(influxdbTemplate))
	},
}

//go:embed tpls/protoc.tpl
var protocTemplate string

var cmdGenProtoc = &cli.Command{
	Name:  "protoc",
	Usage: "generate protoc install script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "protoc.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(protocTemplate))
	},
}

//go:embed tpls/git.tpl
var genGitTemplate string

var cmdGenGit = &cli.Command{
	Name:  "git",
	Usage: "generate git install script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "git.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(genGitTemplate))
	},
}

//go:embed tpls/golang.tpl
var genGolangTemplate string

var cmdGenGolang = &cli.Command{
	Name:  "golang",
	Usage: "generate golang install script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "golang.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(genGolangTemplate))
	},
}

//go:embed tpls/dsh.tpl
var genDshTemplate string

var cmdGenDsh = &cli.Command{
	Name:  "dsh",
	Usage: "generate deepseek harness install script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "dsh.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(genDshTemplate))
	},
}

//go:embed tpls/nvm.tpl
var nvmTemplate string

var cmdGenNvm = &cli.Command{
	Name:  "nvm",
	Usage: "generate nvm script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_Name,
			Aliases: []string{"n"},
			Usage:   "script output file name",
			Value:   "nvm.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		return generateFile(ctx.String(cmdFlag_Output), ctx.String(cmdFlag_Name), []byte(nvmTemplate))
	},
}

//go:embed tpls/k3s-master.tpl
var k3sMasterTemplate string

//go:embed tpls/k3s-worker.tpl
var k3sWorkerTemplate string

var cmdGenK3S = &cli.Command{
	Name:  "k3s",
	Usage: "generate k3s master & worker script",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Output,
			Aliases: []string{"o"},
			Usage:   "script output directory",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "master",
			Aliases: []string{"m"},
			Usage:   "k3s master output file name",
			Value:   "k3s-master.sh",
		},
		&cli.StringFlag{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "k3s worker output file name",
			Value:   "k3s-worker.sh",
		},
	},
	Action: func(ctx *cli.Context) error {
		err := generateFile(ctx.String(cmdFlag_Output), ctx.String("master"), []byte(k3sMasterTemplate))
		if err != nil {
			return err
		}
		return generateFile(ctx.String(cmdFlag_Output), ctx.String("worker"), []byte(k3sWorkerTemplate))
	},
}
