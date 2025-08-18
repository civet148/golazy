package {{.pkgName}}

import (
	{{.imports}}
)

type {{.logic}} struct {
    ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

{{if .hasDoc}}{{.doc}}{{end}}
func New{{.logic}}(c *gin.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
        ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}(ctx context.Context{{if .HasRequest}}, {{.request}}{{end}}) {{.responseType}} {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	{{.returnString}}
}

