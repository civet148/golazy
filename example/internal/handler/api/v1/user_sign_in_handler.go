package v1

import (
	"example/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"

	"example/internal/logic/api/v1"
	"example/internal/svc"
	"github.com/civet148/log"
)

// @Summary 用户登录
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param UserSignIn body types.UserSignInReq true "params description"
// @Success 200 {object} types.UserSignInRsp
// @Router /api/v1/sign_in [post]
func UserSignInHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignInReq
		if err := c.ShouldBind(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := v1.NewUserSignInLogic(c, svcCtx)

		resp, err := l.UserSignIn(c, &req)
		if err != nil {
			log.Errorf("call UserSignIn failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
