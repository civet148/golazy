package gen

import (
	_ "embed"
)

//go:embed tpls/gomod.tpl
var gomodTemplate string

func genGoMod(cfg *Config) error {
	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          "",
		filename:        "go.mod",
		templateName:    "gomodTemplate",
		category:        category,
		builtinTemplate: gomodTemplate,
		data: map[string]string{
			"serviceName": defaultName,
		},
	})
}
