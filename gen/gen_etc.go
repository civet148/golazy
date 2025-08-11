package gen

import (
	_ "embed"
	"github.com/civet148/golazy/parser"
	"path"
	"strconv"
	"strings"
)

var (
	defaultName = "api-server"
)

const (
	defaultPort = 8888
	etcDir      = ""
)

//go:embed tpls/etc.tpl
var etcTemplate string

func genEtc(cfg *Config, api *parser.ApiService) error {
	host := "0.0.0.0"
	port := strconv.Itoa(defaultPort)

	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          etcDir,
		filename:        "config.yaml",
		templateName:    "etcTemplate",
		category:        category,
		builtinTemplate: etcTemplate,
		data: map[string]string{
			"serviceName": defaultName,
			"host":        host,
			"port":        port,
		},
	})
}

func setDefaultName(filename string) {
	//将文件路径前缀和.api文件后缀去掉，取中间的名字作为程序的默认名字
	basePath := path.Base(filename)
	defaultName = strings.Replace(basePath, ".api", "", -1)
}
