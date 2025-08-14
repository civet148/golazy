package {{.pkgName}}

import (
	{{.imports}}
)

type {{.logic}} struct {
    ctx context.Context
	svcCtx *svc.ServiceContext
}

{{if .hasDoc}}{{.doc}}{{end}}
func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
        ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}(ctx context.Context{{if .HasRequest}}, {{.request}}{{end}}) {{.responseType}} {
	// todo: add your logic here and delete this line
	// you can call ctx.(*gin.Context) convert to gin context
	{{.returnString}}
}

