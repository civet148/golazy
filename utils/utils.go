package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/civet148/golazy/parser"
)

func GetMiddleware(as *parser.ApiServer) []string {
	var results []string
	middleware := as.Middleware
	if len(middleware) > 0 {
		for _, item := range strings.Split(middleware, ",") {
			if item != "" {
				results = append(results, strings.TrimSpace(item))
			}
		}
	}
	return results
}

func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil //文件或目录存在
	}
	if os.IsNotExist(err) {
		return false, nil // 文件或目录不存在
	}
	return false, err // 其他错误（如权限问题）
}

func ShellExec(shell string) (err error) {
	fmt.Println(shell)
	cmd := exec.Command("/bin/sh", "-c", shell)
	// 获取标准输出和错误输出
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 执行命令
	err = cmd.Run()
	// 打印输出信息
	if stdout.Len() > 0 {
		fmt.Printf("\n%s", stdout.String())
	}
	if stderr.Len() > 0 {
		fmt.Printf("\n%s", stderr.String())
	}
	if err != nil {
		fmt.Printf("error:%s\n", err)
	}
	return nil
}
