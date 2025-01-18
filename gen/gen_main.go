package gen

import (
	_ "embed"
	"fmt"
	"github.com/civet148/golazy/utils"
	"strings"
	"time"
)

/*-----------------------------------------------------------------------------------------------------------*/

//go:embed tpls/main.tpl
var mainTemplate string

func genMain(cfg *Config, rootPkg string) error {
	name := strings.ToLower(defaultName)
	filename, err := utils.FileNamingFormat(cfg.Style, name)
	if err != nil {
		return err
	}

	configName := filename
	if strings.HasSuffix(filename, "-api") {
		filename = strings.ReplaceAll(filename, "-api", "")
	}

	return genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          "",
		filename:        filename + ".go",
		templateName:    "mainTemplate",
		category:        category,
		builtinTemplate: mainTemplate,
		data: map[string]string{
			"importPackages": genMainImports(rootPkg),
			"serviceName":    configName,
			"datetime":       time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}

func genMainImports(parentPkg string) string {
	var imports []string
	imports = append(imports, fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, configDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, handlerDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"\n", utils.JoinPackages(parentPkg, contextDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"", ProjectGinURL))
	imports = append(imports, fmt.Sprintf("\"%s\"", ProjectLogURL))
	return strings.Join(imports, "\n\t")
}

