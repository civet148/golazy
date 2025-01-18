package {{.PkgName}}

import (
    "context"
	"net/http"
    "github.com/civet148/log"
	"github.com/gin-gonic/gin"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := {{.shouldBind}}; err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
        log.Json("request", req)
		{{end}}l := {{.LogicName}}.New{{.LogicType}}(context.Background(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))
	}
}

