package gen

import (
	_ "embed"
	"fmt"
)

//go:embed tpls/docs.tpl
var docsTemplate string

func genDocs(cfg *Config) error {
	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          "docs",
		filename:        "docs.go",
		templateName:    "docsTemplate",
		category:        category,
		builtinTemplate: docsTemplate,
		data: map[string]string{
			"serviceName": defaultName,
			"hostPort":    getLocalHostPort(),
		},
	})
}

func getLocalHostPort() string {
	return fmt.Sprintf("localhost:%v", defaultPort)
}

func getListenHostPort() string {
	return fmt.Sprintf("0.0.0.0:%v", defaultPort)
}
