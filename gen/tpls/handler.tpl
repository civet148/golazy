package {{.PkgName}}

import (
	"github.com/gin-gonic/gin"
	{{.ImportPackages}}
)

// @Summary {{.Doc}}
// @Description
// @Tags
// @Accept {{.Accept}}
// @Produce {{.Produce}}
// @Param {{.Call}} body {{if .HasRequest}}types.{{.RequestType}}{{else}}string{{end}} true "request params description"
// @Success 200 {{.Object}} {{if .HasRequest}}types.{{.ResponseType}}{{else}}string{{end}}
// @Router {{.RouterPath}} [{{.Method}}]
func {{.HandlerName}}(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
	 {{if .IsNormal}}
	    {{if .HasRequest}}var req types.{{.RequestType}}
		if err := {{.ShouldBind}}; err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
        log.Debugf("request [%+v]", req)
		{{end}}l := {{.LogicName}}.New{{.LogicType}}(c, svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}(c, {{if .HasRequest}}&req{{end}})
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))
		{{else}}
        l := {{.LogicName}}.New{{.LogicType}}(c, svcCtx)
		_ = l.{{.Call}}(c)
	 {{end}}
	}
}

