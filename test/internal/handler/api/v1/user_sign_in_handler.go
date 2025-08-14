package v1

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 用户登录
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param UserSignIn body types.UserSignInReq true "request params description"
// @Success 200 {object} types.UserSignInRsp
// @Router /api/v1/sign_in [post]
func UserSignInHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignInReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := v1.NewUserSignInLogic(c, svcCtx)
		resp, err := l.UserSignIn(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
