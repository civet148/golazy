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

// 用户退出登录
func UserSignOutHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignOutReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := v1.NewUserSignOutLogic(context.Background(), svcCtx)
		resp, err := l.UserSignOut(&req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
