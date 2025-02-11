package gen

import (
	"bytes"
	_ "embed"
	"github.com/civet148/golazy/parser"
	"github.com/civet148/golazy/utils"
	"text/template"
)

const (
	category                    = "api"
	configTemplateFile          = "config.tpl"
	contextTemplateFile         = "context.tpl"
	etcTemplateFile             = "etc.tpl"
	handlerTemplateFile         = "handler.tpl"
	logicTemplateFile           = "logic.tpl"
	mainTemplateFile            = "main.tpl"
	middlewareImplementCodeFile = "middleware.tpl"
	routesTemplateFile          = "routes.tpl"
	routesAdditionTemplateFile  = "route-addition.tpl"
	typesTemplateFile           = "types.tpl"
)

var templates = map[string]string{
	//configTemplateFile:          configTemplate,
	contextTemplateFile: contextTemplate,
	//etcTemplateFile:             etcTemplate,
	//handlerTemplateFile:         handlerTemplate,
	//logicTemplateFile:           logicTemplate,
	//mainTemplateFile:            mainTemplate,
	//middlewareImplementCodeFile: middlewareImplementCode,
	//routesTemplateFile:          routesTemplate,
	//routesAdditionTemplateFile:  routesAdditionTemplate,
	//typesTemplateFile:           typesTemplate,
}

type fileGenConfig struct {
	dir             string
	subdir          string
	filename        string
	templateName    string
	category        string
	builtinTemplate string
	data            any
}

type Config struct {
	ApiFile string
	OutDir  string
	Style   string
}

func GenerateGoCode(cfg *Config, services []*parser.ApiService) (err error) {

	setDefaultName(cfg.ApiFile)

	// 打印解析结果
	utils.Must(utils.MkdirIfNotExist(cfg.OutDir))
	rootPkg, err := utils.GetParentPackage(cfg.OutDir)
	if err != nil {
		return err
	}

	utils.Must(genMain(cfg, rootPkg))
	for _, svc := range services {
		utils.Must(genEtc(cfg, svc))
		utils.Must(genConfig(cfg, svc))
		utils.Must(genHandler(cfg, rootPkg, svc))
		utils.Must(genLogic(cfg, rootPkg, svc))
		utils.Must(genMiddleware(cfg, svc))
		utils.Must(genTypes(cfg, svc))
		utils.Must(genServiceContext(cfg, rootPkg))
	}
	utils.Must(genRoutes(cfg, rootPkg, services))
	utils.Must(genScript(cfg, rootPkg))
	utils.Must(genMakefile(cfg, rootPkg))
	utils.Must(genDockerfile(cfg, rootPkg))
	return nil
}

func genFile(c fileGenConfig) error {
	fp, created, err := utils.MaybeCreateFile(c.dir, c.subdir, c.filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	t := template.Must(template.New(c.templateName).Parse(c.builtinTemplate))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.data)
	if err != nil {
		return err
	}

	code := utils.FormatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

