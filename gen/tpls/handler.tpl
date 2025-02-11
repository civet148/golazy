package {{.PkgName}}

import (
    "context"
	"github.com/gin-gonic/gin"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
	 {{if .IsNormal}}
	    {{if .HasRequest}}var req types.{{.RequestType}}
		if err := {{.shouldBind}}; err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
        log.Debugf("request [%+v]", req)
		{{end}}l := {{.LogicName}}.New{{.LogicType}}(context.Background(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))
		{{else}}
        l := {{.LogicName}}.New{{.LogicType}}(context.Background(), svcCtx)
		_ = l.{{.Call}}(c)
	 {{end}}
	}
}

