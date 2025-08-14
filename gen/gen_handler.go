package gen

import (
	_ "embed"
	"fmt"
	"github.com/civet148/golazy/parser"
	"github.com/civet148/golazy/utils"
	"github.com/civet148/log"
	"path"
	"strings"
)

const defaultLogicPackage = "logic"

//go:embed tpls/handler.tpl
var handlerTemplate string

func genHandler(cfg *Config, rootPkg string, api *parser.ApiService) error {

	var err error
	for _, spec := range api.APIs {
		if spec.Request == typesGinContext {
			err = genContextHandler(cfg, rootPkg, api, spec) //generate context handler
			if err != nil {
				return log.Errorf(err.Error())
			}
		} else {
			err = genNormalHandler(cfg, rootPkg, api, spec) //generate normal handler
			if err != nil {
				return log.Errorf(err.Error())
			}
		}
	}
	return nil
}

func genNormalHandler(cfg *Config, rootPkg string, api *parser.ApiService, spec *parser.ApiSpec) (err error) {

	handler := getHandlerName(spec.Handler)
	handlerPath := getHandlerFolderPath(api.Server.Group, api.Server.Prefix)
	pkgName := handlerPath[strings.LastIndex(handlerPath, "/")+1:]
	logicPkgName := defaultLogicPackage
	if handlerPath != handlerDir {
		handler = strings.Title(handler)
		logicPkgName = pkgName
	}
	filename, err := utils.FileNamingFormat(cfg.Style, handler)
	if err != nil {
		return err
	}
	var strShouldBind = "c.ShouldBind(&req)"
	//路由中包含变量或正则表达式
	if strings.Contains(spec.Path, ":") || strings.Contains(spec.Path, "*") || strings.Contains(spec.Path, "{") {
		strShouldBind = "c.ShouldBindUri(&req)"
	}
	var routerPath string
	routerPath = getRouterPath(api, spec)
	err = genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          getHandlerFolderPath(api.Server.Group, api.Server.Prefix),
		filename:        filename + ".go",
		templateName:    "handlerTemplate",
		category:        category,
		builtinTemplate: handlerTemplate,
		data: map[string]any{
			"PkgName":        pkgName,
			"IsNormal":       true,
			"ImportPackages": getNormalHandlerImports(api.Server.Group, api.Server.Prefix, rootPkg),
			"HandlerName":    handler,
			"RequestType":    spec.Request,
			"ResponseType":   spec.Response,
			"LogicName":      logicPkgName,
			"LogicType":      strings.Title(getLogicName(handler)),
			"Call":           strings.Title(strings.TrimSuffix(handler, "Handler")),
			"HasResp":        canGenTypes(spec.Response),
			"HasRequest":     len(spec.Request) > 0,
			"HasDoc":         len(spec.Doc) > 0,
			"Doc":            spec.Doc,
			"ShouldBind":     strShouldBind,
			"Accept":         "json",
			"Produce":        "json",
			"RouterPath":     routerPath,
			"Method":         spec.Method,
			"Object":         "{object}",
		},
	})
	if err != nil {
		return log.Errorf(err.Error())
	}
	return nil
}

func genContextHandler(cfg *Config, rootPkg string, api *parser.ApiService, spec *parser.ApiSpec) (err error) {
	handler := getHandlerName(spec.Handler)
	handlerPath := getHandlerFolderPath(api.Server.Group, api.Server.Prefix)
	pkgName := handlerPath[strings.LastIndex(handlerPath, "/")+1:]
	logicPkgName := defaultLogicPackage
	if handlerPath != handlerDir {
		handler = strings.Title(handler)
		logicPkgName = pkgName
	}
	filename, err := utils.FileNamingFormat(cfg.Style, handler)
	if err != nil {
		return err
	}
	var routerPath string
	routerPath = getRouterPath(api, spec)
	err = genFile(fileGenConfig{
		dir:             cfg.OutDir,
		subdir:          getHandlerFolderPath(api.Server.Group, api.Server.Prefix),
		filename:        filename + ".go",
		templateName:    "handlerTemplate",
		category:        category,
		builtinTemplate: handlerTemplate,
		data: map[string]any{
			"PkgName":        pkgName,
			"IsNormal":       false,
			"ImportPackages": getContextHandlerImports(api.Server.Group, api.Server.Prefix, rootPkg),
			"HandlerName":    handler,
			"LogicName":      logicPkgName,
			"LogicType":      strings.Title(getLogicName(handler)),
			"Call":           strings.Title(strings.TrimSuffix(handler, "Handler")),
			"HasResp":        false,
			"HasRequest":     false,
			"HasDoc":         len(spec.Doc) > 0,
			"Doc":            spec.Doc,
			"Accept":         "plain",
			"Produce":        "plain",
			"RouterPath":     routerPath,
			"Method":         spec.Method,
			"Object":         "{string}",
		},
	})
	if err != nil {
		return log.Errorf(err.Error())
	}
	return nil
}

func getCommentDoc(doc string) string {
	return fmt.Sprintf("//%s", doc)
}

func getLogicFolderPath(group, route string) string {
	folder := route
	if len(folder) == 0 {
		if len(folder) == 0 {
			return logicDir
		}
	}
	folder = strings.TrimPrefix(folder, "/")
	folder = strings.TrimSuffix(folder, "/")
	return path.Join(logicDir, folder)
}

func getNormalHandlerImports(group, route string, parentPkg string) string {
	imports := []string{
		fmt.Sprintf("\"%s\"", "net/http"),
		fmt.Sprintf("\"%s\"", "github.com/civet148/log"),
		fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, getLogicFolderPath(group, route))),
		fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, contextDir)),
	}
	if len(route) > 0 {
		imports = append(imports, fmt.Sprintf("\"%s\"\n", utils.JoinPackages(parentPkg, typesDir)))
	}

	return strings.Join(imports, "\n\t")
}

func getContextHandlerImports(group, route string, parentPkg string) string {
	imports := []string{
		fmt.Sprintf("\"%s\"", "github.com/civet148/log"),
		fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, getLogicFolderPath(group, route))),
		fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, contextDir)),
	}
	return strings.Join(imports, "\n\t")
}

func getHandlerBaseName(handler string) (string, error) {
	handler = strings.TrimSpace(handler)
	handler = strings.TrimSuffix(handler, "handler")
	handler = strings.TrimSuffix(handler, "Handler")
	return handler, nil
}

func getHandlerFolderPath(group, prefix string) string {
	folder := prefix
	if len(folder) == 0 {
		return handlerDir
	}
	folder = strings.TrimPrefix(folder, "/")
	folder = strings.TrimSuffix(folder, "/")
	return path.Join(handlerDir, folder)
}

func getHandlerName(route string) string {
	handler, err := getHandlerBaseName(route)
	if err != nil {
		panic(err)
	}

	return handler + "Handler"
}

func getLogicName(route string) string {
	handler, err := getHandlerBaseName(route)
	if err != nil {
		panic(err)
	}

	return handler + "Logic"
}
