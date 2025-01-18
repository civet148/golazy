package gen

import (
	_ "embed"
	"github.com/civet148/golazy/utils"
)

/*-----------------------------------------------------------------------------------------------------------*/

const contextFilename = "service_context"

//go:embed tpls/svc.tpl
var contextTemplate string

func genServiceContext(cfg *Config, rootPkg string) error {
	filename, err := utils.FileNamingFormat(cfg.Style, contextFilename)
	if err != nil {
		return err
	}

	var middlewareStr string
	var middlewareAssignment string

	configImport := "\"" + utils.JoinPackages(rootPkg, configDir) + "\""
	if len(middlewareStr) > 0 {
		configImport += "\n\t\"" + utils.JoinPackages(rootPkg, middlewareDir) + "\""
	}

	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          contextDir,
		filename:        filename + ".go",
		templateName:    "contextTemplate",
		category:        category,
		builtinTemplate: contextTemplate,
		data: map[string]string{
			"configImport":         configImport,
			"config":               "config.Config",
			"middleware":           middlewareStr,
			"middlewareAssignment": middlewareAssignment,
		},
	})
}

