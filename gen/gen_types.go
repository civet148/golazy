package gen

import (
	_ "embed"
	"fmt"
	"github.com/civet148/golazy/parser"
	"github.com/civet148/golazy/utils"
	"github.com/civet148/log"
	"io"
	"strings"
)

const (
	typesNil        = "nil"
	typesGinContext = "gin.Context"
)

//go:embed tpls/types.tpl
var typesTemplate string

var specialTypes = []string{
	typesNil,
	typesGinContext,
}

func canGenTypes(name string) bool {
	for _, v := range specialTypes {
		if name == v {
			return false
		}
	}
	if strings.Contains(name, ".") {
		return false
	}
	return true
}

func genTypes(cfg *Config, api *parser.ApiService) error {
	for _, spec := range api.APIs {
		typeReqVal := buildTypes([]string{spec.Request})
		typeReqFile, err := utils.FileNamingFormat(cfg.Style, spec.Request)
		if err != nil {
			return log.Errorf(err.Error())
		}
		if canGenTypes(spec.Request) {
			err = createTypesFile(cfg.OutDir, typeReqVal, typeReqFile)
			if err != nil {
				return log.Errorf(err.Error())
			}
		}

		typeRspVal := buildTypes([]string{spec.Response})
		typeRspFile, err := utils.FileNamingFormat(cfg.Style, spec.Response)
		if err != nil {
			return log.Errorf(err.Error())
		}
		if canGenTypes(spec.Response) {
			err = createTypesFile(cfg.OutDir, typeRspVal, typeRspFile)
			if err != nil {
				return log.Errorf(err.Error())
			}
		}
	}
	return nil
}

func createTypesFile(dir, typeValue, typeFilename string) (err error) {

	if !strings.HasSuffix(typeFilename, ".go") {
		typeFilename = typeFilename + ".go"
	}

	err = genFile(fileGenConfig{
		dir:             dir,
		subdir:          typesDir,
		filename:        typeFilename,
		templateName:    "typesTemplate",
		category:        category,
		builtinTemplate: typesTemplate,
		data: map[string]any{
			"types":        typeValue,
			"containsTime": false,
		},
	})
	if err != nil {
		return log.Errorf(err.Error())
	}
	return nil
}

// buildTypes gen types to string
func buildTypes(types []string) string {
	var builder strings.Builder
	first := true
	for _, tp := range types {
		if first {
			first = false
		} else {
			builder.WriteString("\n\n")
		}
		writeType(&builder, tp)
	}
	return builder.String()
}

func writeType(writer io.Writer, tp string) {
	fmt.Fprintf(writer, "type %s struct {\n", tp)
	fmt.Fprintf(writer, "}")
}

