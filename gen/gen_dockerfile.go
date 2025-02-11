package gen

import (
	_ "embed"
	"strings"
)

/*-----------------------------------------------------------------------------------------------------------*/

//go:embed tpls/dockerfile.tpl
var dockerTemplate string

func genDockerfile(cfg *Config, rootPkg string) error {
	name := strings.ToLower(defaultName)
	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          "",
		filename:        "Dockerfile",
		templateName:    "dockerTemplate",
		category:        category,
		builtinTemplate: dockerTemplate,
		data: map[string]string{
			"ProgramName": name,
		},
	})
}

