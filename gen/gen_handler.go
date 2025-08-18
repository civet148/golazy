package gen

import (
	_ "embed"
	"fmt"
	"github.com/civet148/golazy/parser"
	"github.com/civet148/golazy/utils"
	"github.com/civet148/log"
	"path"
	"regexp"
	"strings"
)

const defaultLogicPackage = "logic"

//go:embed tpls/handler.tpl
var handlerTemplate string

func genHandlers(cfg *Config, rootPkg string, api *parser.ApiService) error {

	var err error
	for _, spec := range api.APIs {
		err = genApiHandler(cfg, rootPkg, api, spec) //generate normal handler
		if err != nil {
			return log.Errorf(err.Error())
		}
	}
	return nil
}

func genApiHandler(cfg *Config, rootPkg string, api *parser.ApiService, spec *parser.ApiSpec) (err error) {

	handler := getHandlerName(spec.Handler)
	handlerPath := getHandlerFolderPath(api.Server.Group, api.Server.Prefix)
	pkgName := handlerPath[strings.LastIndex(handlerPath, "/")+1:]
	logicPkgName := defaultLogicPackage
	if handlerPath != handlerDir {
		handler = strings.Title(handler)
		logicPkgName = pkgName
	}
	hasReq := canGenTypes(spec.Request)
	hasResp := canGenTypes(spec.Response)

	filename, err := utils.FileNamingFormat(cfg.Style, handler)
	if err != nil {
		return err
	}
	var strShouldBind = "c.ShouldBind(&req)"
	//路由中包含变量或正则表达式
	if isRegexpRoute(spec.Path) {
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
			"ImportPackages": getHandlerImports(api.Server.Group, api.Server.Prefix, rootPkg, hasReq || hasResp),
			"HandlerName":    handler,
			"RequestType":    spec.Request,
			"ResponseType":   spec.Response,
			"LogicName":      logicPkgName,
			"LogicType":      strings.Title(getLogicName(handler)),
			"Call":           strings.Title(strings.TrimSuffix(handler, "Handler")),
			"HasRequest":     hasReq,
			"HasResp":        hasResp,
			"HasDoc":         len(spec.Doc) > 0,
			"Doc":            spec.Doc,
			"ShouldBind":     strShouldBind,
			"Accept":         "json",
			"Produce":        "json",
			"RouterPath":     simplifyRouteForSwag(routerPath),
			"Method":         spec.Method,
			"Object":         "{object}",
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

func isRegexpRoute(path string) bool {
	return strings.Contains(path, ":") || strings.Contains(path, "*") ||
		strings.Contains(path, "{") || strings.Contains(path, "(")
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

func getHandlerImports(group, route string, parentPkg string, hasTypes bool) string {
	var imports []string
	if hasTypes {
		imports = append(imports, fmt.Sprintf("\"%s\"", "net/http"))
		imports = append(imports, fmt.Sprintf("\"%s\"\n", utils.JoinPackages(parentPkg, typesDir)))
	}

	imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/civet148/log"))
	imports = append(imports, fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, getLogicFolderPath(group, route))))
	imports = append(imports, fmt.Sprintf("\"%s\"", utils.JoinPackages(parentPkg, contextDir)))
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

// 移除路由中的正则表达式（用于swag注释）
func simplifyRouteForSwag(route string) string {
	// 处理 /path/{param:[0-9]+} 格式
	re := regexp.MustCompile(`\{(\w+):[^}]+\}`)
	simplified := re.ReplaceAllString(route, `:$1`)

	// 处理 /path/:param([0-9]+) 格式
	re = regexp.MustCompile(`:(\w+)\([^)]+\)`)
	simplified = re.ReplaceAllString(simplified, `:$1`)

	return simplified
}
