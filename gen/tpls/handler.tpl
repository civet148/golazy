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
// @Param {{.Call}} body {{if .HasRequest}}types.{{.RequestType}}{{else}}string{{end}} true "params description"
// @Success 200 {{.Object}} {{if .HasResp}}types.{{.ResponseType}}{{else}}nil{{end}}
// @Router {{.RouterPath}} [{{.Method}}]
func {{.HandlerName}}(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
	 {{if .IsNormal}}
	    {{if .HasRequest}}
	    var req types.{{.RequestType}}
		if err := {{.ShouldBind}}; err != nil {
			if err != nil {
                log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
            }
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
        log.Infof("request: %+v", req)
		{{end}}
		l := {{.LogicName}}.New{{.LogicType}}(c, svcCtx)
		{{if .HasResp}}
		resp, err := l.{{.Call}}(c, {{if .HasRequest}}&req{{end}})
		if err != nil {
			log.Errorf("call {{.Call}} failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))
		{{else}}
		err := l.{{.Call}}(c, {{if .HasRequest}}&req{{end}})
		if err != nil {
            log.Errorf("call {{.Call}} failed, err: %v", err.Error())
        }
        c.Abort()
		{{end}}
    {{else}}
        l := {{.LogicName}}.New{{.LogicType}}(c, svcCtx)
		err := l.{{.Call}}(c)
		if err != nil {
			log.Errorf("call {{.Call}} failed, err: %v", err.Error())
		}
        c.Abort()
	{{end}}
	}
}

