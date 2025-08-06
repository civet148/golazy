package user

import (
	"context"
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/v1/user"
	"test/internal/svc"
	"test/internal/types"
)

// 修改用户
func EditUserHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.EditUserReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := user.NewEditUserLogic(context.Background(), svcCtx)
		resp, err := l.EditUser(&req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
