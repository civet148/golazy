package gen

import (
	_ "embed"
	"github.com/civet148/golazy/utils"
	"runtime"
	"strings"
)

const (
	scriptName = "genModels"
)

/*-----------------------------------------------------------------------------------------------------------*/

//go:embed tpls/db2go_sh.tpl
var db2goShellTemplate string

//go:embed tpls/db2go_bat.tpl
var db2goBatchTemplate string

func genScript(cfg *Config, rootPkg string) error {

	filename, err := utils.FileNamingFormat(cfg.Style, scriptName)
	if err != nil {
		return err
	}

	var strTemplate string
	var ext string
	switch runtime.GOOS {
	case OsWindows:
		ext = ".bat"
		strTemplate = db2goBatchTemplate
	default:
		ext = ".sh"
		strTemplate = db2goShellTemplate
	}
	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          internal,
		filename:        filename + ext,
		templateName:    "scriptTemplate",
		category:        category,
		builtinTemplate: strTemplate,
		data: map[string]string{
			"importModel": genImportModel(rootPkg),
		},
	})
}

func genImportModel(parentPkg string) string {
	return strings.TrimSuffix(parentPkg+"/"+internal, "/")
}

