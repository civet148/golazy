package v1

import (
	"example/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"

	"example/internal/logic/api/v1"
	"example/internal/svc"
	"github.com/civet148/log"
)

// @Summary 用户退出登录
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param UserSignOut body types.UserSignOutReq true "params description"
// @Success 200 {object} types.UserSignOutRsp
// @Router /api/v1/sign_out [post]
func UserSignOutHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignOutReq
		if err := c.ShouldBind(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := v1.NewUserSignOutLogic(c, svcCtx)

		resp, err := l.UserSignOut(c, &req)
		if err != nil {
			log.Errorf("call UserSignOut failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
