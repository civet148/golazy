package middleware

import (
    "github.com/gin-gonic/gin"
    {{if .hasNetHttp}}"net/http"{{end}}
    {{.importPkg}}
)

{{.constants}}

type {{.name}} struct {
    {{if .hasWhiteList}}WhiteList map[string]bool{{end}}
}

func New{{.name}}() *{{.name}} {
	return &{{.name}}{
	{{if .hasWhiteList}}WhiteList: map[string]bool{}, {{end}}
	}
}

func (m *{{.name}}) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
        {{.defaultMiddlewareImpl}}

        //TODO: add your middleware logic here

		// Pass through to next handler
		c.Next()
	}
}

{{.extraFunctions}}
