package gen

import (
	_ "embed"
	"github.com/civet148/golazy/utils"
	"github.com/civet148/log"
	"os/exec"
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
	err = genFile(fileGenConfig{
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
	if err != nil {
		return log.Errorf(err.Error())
	}
	var strScript = cfg.OutDir + "/" + internal + "/" + filename + ext
	switch runtime.GOOS {
	case OsWindows:
	default:
		err = exec.Command("chmod", "+x", strScript).Run()
		if err != nil {
			return log.Errorf("chmod error: %s", err.Error())
		}
	}
	return nil
}

func genImportModel(parentPkg string) string {
	return strings.TrimSuffix(parentPkg+"/"+internal, "/")
}
