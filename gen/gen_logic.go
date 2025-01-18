package gen

import (
	_ "embed"
	"fmt"
	"github.com/civet148/golazy/parser"
	"github.com/civet148/golazy/utils"
	"github.com/civet148/log"
	"strings"
)

//go:embed tpls/logic.tpl
var logicTemplate string

func genLogic(cfg *Config, rootPkg string, api *parser.ApiService) error {
	err := genLogicByRoute(cfg, rootPkg, api)
	if err != nil {
		return err
	}
	return nil
}

func genLogicByRoute(cfg *Config, rootPkg string, api *parser.ApiService) error {
	for _, spec := range api.APIs {
		logic := getLogicName(spec.Handler)
		goFile, err := utils.FileNamingFormat(cfg.Style, logic)
		if err != nil {
			return err
		}

		imports := genLogicImports(rootPkg)
		var responseString string
		var returnString string
		var requestString string
		if len(spec.Response) > 0 {
			resp := responseGoTypeName(spec.Response, typesPacket)
			responseString = "(resp " + resp + ", err error)"
			returnString = "return"
		} else {
			responseString = "error"
			returnString = "return nil"
		}
		if len(spec.Request) > 0 {
			requestString = "req *" + requestGoTypeName(spec.Request, typesPacket)
		}

		subDir := getLogicFolderPath(api.Server.Group, api.Server.Prefix)
		err = genFile(fileGenConfig{
			dir:             cfg.OutDir,
			subdir:          subDir,
			filename:        goFile + ".go",
			templateName:    "logicTemplate",
			category:        category,
			builtinTemplate: logicTemplate,
			data: map[string]any{
				"pkgName":      subDir[strings.LastIndex(subDir, "/")+1:],
				"imports":      imports,
				"logic":        strings.Title(logic),
				"function":     strings.Title(strings.TrimSuffix(logic, "Logic")),
				"responseType": responseString,
				"returnString": returnString,
				"request":      requestString,
				"hasDoc":       len(spec.Doc) > 0,
				"doc":          getDoc(spec.Doc),
			},
		})
		if err != nil {
			return log.Errorf(err.Error())
		}
	}
	return nil
}

func genLogicImports(parentPkg string) string {
	var imports []string
	imports = append(imports, `"context"`+"\n")
	imports = append(imports, fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, contextDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"\n", utils.JoinPackages(parentPkg, typesDir)))
	return strings.Join(imports, "\n\t")
}

func responseGoTypeName(resp, pkg string) string {
	if strings.HasPrefix(resp, "*") {
		resp = strings.Replace(resp, "*", "", -1)
	}
	return "*" + pkg + "." + resp
}

func requestGoTypeName(req, pkg string) string {
	if strings.HasPrefix(req, "*") {
		req = strings.Replace(req, "*", "", -1)
	}
	return pkg + "." + req
}

