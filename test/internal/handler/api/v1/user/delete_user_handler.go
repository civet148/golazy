package user

import (
	"context"
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1/user"
	"test/internal/svc"
	"test/internal/types"
)

// 删除用户
func DeleteUserHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.DeleteUserReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := user.NewDeleteUserLogic(context.Background(), svcCtx)
		resp, err := l.DeleteUser(&req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
