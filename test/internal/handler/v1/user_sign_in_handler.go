package v1

import (
	"context"
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/v1"
	"test/internal/svc"
	"test/internal/types"
)

// 用户登录
func UserSignInHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignInReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := v1.NewUserSignInLogic(context.Background(), svcCtx)
		resp, err := l.UserSignIn(&req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
