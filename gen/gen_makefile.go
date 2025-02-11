package gen

import (
	_ "embed"
	"strings"
)

/*-----------------------------------------------------------------------------------------------------------*/

//go:embed tpls/makefile.tpl
var makefileTemplate string

func genMakefile(cfg *Config, rootPkg string) error {
	name := strings.ToLower(defaultName)
	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          "",
		filename:        "Makefile",
		templateName:    "mainTemplate",
		category:        category,
		builtinTemplate: makefileTemplate,
		data: map[string]string{
			"ProgramName": name,
		},
	})
}

