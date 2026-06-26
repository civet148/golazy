package utils

import (
	"os"
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
