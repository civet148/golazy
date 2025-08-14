package v1

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 用户退出登录
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param UserSignOut body types.UserSignOutReq true "request params description"
// @Success 200 {object} types.UserSignOutRsp
// @Router /api/v1/sign_out [post]
func UserSignOutHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignOutReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := v1.NewUserSignOutLogic(c, svcCtx)
		resp, err := l.UserSignOut(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
