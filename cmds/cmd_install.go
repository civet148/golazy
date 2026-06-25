package cmds

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
)

const (
	Cmd_Install = "install"
)

const (
	subCmd_InstallGrpcGateway = "grpc-gateway"
	subCmd_InstallDB2GO       = "db2go"
)

const (
	cmdFlag_ProtocGenGo          = "protoc-gen-go"
	cmdFlag_ProtocGenGoGrpc      = "protoc-gen-go-grpc"
	cmdFlag_ProtocGenGrpcGateway = "protoc-gen-grpc-gateway"
	cmdFlag_ProtocGenOpenApiV2   = "protoc-gen-openapiv2"
	cmdFlag_GoogleApis           = "google-apis"
	cmdFlag_GogoProtobuf         = "gogo-protobuf"
	cmdFlag_WithSSH              = "with-ssh"
	cmdFlag_Version              = "version"
)

const (
	packageProtocGenGo          = "google.golang.org/protobuf/cmd/protoc-gen-go"
	packageProtocGenGoGrpc      = "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	packageProtocGenGrpcGateway = "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	packageProtocGenOpenApiV2   = "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	packageGoogleApis           = "github.com/googleapis/googleapis"
	packageGogoProtobuf         = "github.com/gogo/protobuf"
	packageDB2GO                = "github.com/civet148/db2go"
)

/*
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.16.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.16.0
*/
var CmdInstall = &cli.Command{
	Name:  Cmd_Install,
	Usage: "install commands",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		cmdInstallGrpcGateway,
		cmdInstallDB2GO,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

var cmdInstallGrpcGateway = &cli.Command{
	Name:  subCmd_InstallGrpcGateway,
	Usage: "install gRPC gateway utils",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_ProtocGenGo,
			Aliases: []string{"g"},
			Usage:   "proto-gen-go version",
			Value:   "v1.28.1",
		},
		&cli.StringFlag{
			Name:    cmdFlag_ProtocGenGoGrpc,
			Aliases: []string{"G"},
			Usage:   "protoc-gen-go-grpc version",
			Value:   "v1.2.0",
		},
		&cli.StringFlag{
			Name:    cmdFlag_ProtocGenGrpcGateway,
			Aliases: []string{"W"},
			Usage:   "protoc-gen-grpc-gateway version",
			Value:   "v2.16.0",
		},
		&cli.StringFlag{
			Name:    cmdFlag_ProtocGenOpenApiV2,
			Aliases: []string{"O"},
			Usage:   "protoc-gen-openapiv2 version",
			Value:   "v2.16.0",
		},
		&cli.StringFlag{
			Name:    cmdFlag_GoogleApis,
			Aliases: []string{"A"},
			Usage:   "github.com/googleapis/googleapis branch",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    cmdFlag_GogoProtobuf,
			Aliases: []string{"P"},
			Usage:   "github.com/gogo/protobuf branch",
			Value:   "",
		},
		&cli.BoolFlag{
			Name:    cmdFlag_WithSSH,
			Aliases: []string{"S"},
			Usage:   "with git ssh to clone",
		},
	},
	Action: func(ctx *cli.Context) error {

		var installPackages = map[string]string{
			packageProtocGenGo:          ctx.String(cmdFlag_ProtocGenGo),
			packageProtocGenGoGrpc:      ctx.String(cmdFlag_ProtocGenGoGrpc),
			packageProtocGenGrpcGateway: ctx.String(cmdFlag_ProtocGenGrpcGateway),
			packageProtocGenOpenApiV2:   ctx.String(cmdFlag_ProtocGenOpenApiV2),
		}
		var clonePackages = map[string]string{
			packageGoogleApis:   ctx.String(cmdFlag_GoogleApis),
			packageGogoProtobuf: ctx.String(cmdFlag_GogoProtobuf),
		}
		// 预定义的常用 protoc-gen 工具列表
		var installPlugins []GoPackageOptions
		for k, v := range installPackages {
			installPlugins = append(installPlugins, GoPackageOptions{
				Package: k,
				Version: v,
			})
		}
		for k, v := range clonePackages {
			installPlugins = append(installPlugins, GoPackageOptions{
				Package: k,
				Version: v,
				Clone:   true,
				WithSSH: ctx.Bool(cmdFlag_WithSSH),
			})
		}
		installer := NewGoInstaller(true)
		return installer.InstallMultiple(installPlugins...)
	},
}

var cmdInstallDB2GO = &cli.Command{
	Name:  subCmd_InstallDB2GO,
	Usage: "install db2go cli",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmdFlag_Version,
			Aliases: []string{"V"},
			Usage:   "db2go version",
			Value:   "latest",
		},
		&cli.BoolFlag{
			Name:    cmdFlag_WithSSH,
			Aliases: []string{"S"},
			Usage:   "with git ssh to clone",
		},
	},
	Action: func(ctx *cli.Context) error {

		var installPackages = map[string]string{
			packageDB2GO: ctx.String(cmdFlag_Version),
		}
		var installPlugins []GoPackageOptions
		for k, v := range installPackages {
			installPlugins = append(installPlugins, GoPackageOptions{
				Package: k,
				Version: v,
			})
		}
		installer := NewGoInstaller(true)
		return installer.InstallMultiple(installPlugins...)
	},
}

// GoPackageOptions 定义 go install 的配置选项
type GoPackageOptions struct {
	// 包名，如 google.golang.org/protobuf/cmd/protoc-gen-go
	Package string
	// 版本号，如 v1.28.1（可选，如果为空则不指定版本）
	Version string
	// 是否在安装前执行 go mod download（可选）
	DownloadFirst bool
	// 工作目录（可选，默认当前目录）
	WorkDir string
	// 是否为克隆模式(默认 go install)
	Clone bool
	// 是否使用git ssh方式
	WithSSH bool
}

// GoInstaller 封装 go install 操作
type GoInstaller struct {
	// 是否打印详细日志
	Verbose bool
}

// NewGoInstaller 创建新的 GoInstaller 实例
func NewGoInstaller(verbose bool) *GoInstaller {
	return &GoInstaller{
		Verbose: verbose,
	}
}

func replaceCloneUrl(packageUrl string, withSSH bool) string {
	if !withSSH {
		return "https://" + packageUrl
	}
	return "git@" + strings.Replace(packageUrl, "/", ":", 1)
}

// Clone 执行git clone指定代码库
func (g *GoInstaller) Clone(opts GoPackageOptions) error {

	// 判断$GOPATH是否存在，如果存在则使用$GOPATH/src目录作为下载基础目录
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return fmt.Errorf("GOPATH environment variable is not set")
	}

	// 创建父级目录
	fullpath := filepath.Join(gopath, "src", opts.Package)
	dir := filepath.Dir(fullpath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("创建目录 %s 失败: %w", dir, err)
	}

	// 构建命令参数(git clone -b v1.2.39 https://github.com/gogo/protobuf.git $GOPATH/src/github.com/gogo/protobuf)
	args := []string{"clone"}
	if opts.Version != "" {
		args = append(args, "-b", opts.Version)
	}
	args = append(args, replaceCloneUrl(opts.Package, opts.WithSSH))
	args = append(args, fullpath)

	// 创建命令
	cmd := exec.Command("git", args...)

	// 设置工作目录
	if opts.WorkDir != "" {
		cmd.Dir = opts.WorkDir
	}

	// 获取标准输出和错误输出
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 打印执行的命令（如果 verbose 模式）
	if g.Verbose {
		log.Printf("执行命令: git %s", strings.Join(args, " "))
		if opts.WorkDir != "" {
			log.Printf("工作目录: %s", opts.WorkDir)
		}
	}

	// 执行命令
	err := cmd.Run()

	// 打印输出信息
	if stdout.Len() > 0 {
		if g.Verbose {
			log.Printf("\n%s", stdout.String())
		}
	}
	if stderr.Len() > 0 {
		if g.Verbose {
			log.Printf("\n%s", stderr.String())
		}
	}

	if err != nil {
		return fmt.Errorf("git clone 执行失败: %w\n标准错误: %s", err, stderr.String())
	}

	if g.Verbose {
		log.Printf("✅ 成功克隆: %s", opts.Package)
	}
	return nil
}

// Install 执行 go install 安装指定的包
func (g *GoInstaller) Install(opts GoPackageOptions) error {
	// 构建完整的包名（带版本号）
	pkgWithVersion := opts.Package
	if opts.Version != "" {
		pkgWithVersion = fmt.Sprintf("%s@%s", opts.Package, opts.Version)
	}

	// 构建命令参数
	args := []string{"install"}
	if opts.DownloadFirst {
		args = append(args, "-mod=mod")
	}
	args = append(args, pkgWithVersion)

	// 创建命令
	cmd := exec.Command("go", args...)

	// 设置工作目录
	if opts.WorkDir != "" {
		cmd.Dir = opts.WorkDir
	}

	// 获取标准输出和错误输出
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 打印执行的命令（如果 verbose 模式）
	if g.Verbose {
		log.Printf("执行命令: go %s", strings.Join(args, " "))
		if opts.WorkDir != "" {
			log.Printf("工作目录: %s", opts.WorkDir)
		}
	}

	// 执行命令
	err := cmd.Run()

	// 打印输出信息
	if stdout.Len() > 0 {
		if g.Verbose {
			log.Printf("\n%s", stdout.String())
		}
	}
	if stderr.Len() > 0 {
		if g.Verbose {
			log.Printf("\n%s", stderr.String())
		}
	}

	if err != nil {
		return fmt.Errorf("go install 执行失败: %w\n标准错误: %s", err, stderr.String())
	}

	if g.Verbose {
		log.Printf("✅ 成功安装: %s", pkgWithVersion)
	}
	return nil
}

// InstallMultiple 批量安装多个包
func (g *GoInstaller) InstallMultiple(packages ...GoPackageOptions) error {
	for _, pkg := range packages {
		if pkg.Clone {
			if err := g.Clone(pkg); err != nil {
				return err
			}
		} else {
			if err := g.Install(pkg); err != nil {
				return err
			}
		}
	}
	return nil
}
