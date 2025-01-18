package gen

import (
	_ "embed"
	"fmt"
	"github.com/civet148/golazy/parser"
	"github.com/civet148/golazy/utils"
)

const (
	configFile = "config"

	jwtTemplate = ` struct {
		AccessSecret string
		AccessExpire int64
	}
`
	jwtTransTemplate = ` struct {
		Secret     string
		PrevSecret string
	}
`
)

//go:embed tpls/config.tpl
var configTemplate string

func genConfig(cfg *Config, api *parser.ApiService) error {
	filename, err := utils.FileNamingFormat(cfg.Style, configFile)
	if err != nil {
		return err
	}

	authImportStr := fmt.Sprintf("\"%s/rest\"", ProjectOpenSourceURL)

	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          configDir,
		filename:        filename + ".go",
		templateName:    "configTemplate",
		category:        category,
		builtinTemplate: configTemplate,
		data: map[string]string{
			"authImport": authImportStr,
		},
	})
}

