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

// 根据ID查询用户
func GetUserByIdHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.GetUserByIdReq
		if err := svc.ShouldBindParams(c, &req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := user.NewGetUserByIdLogic(context.Background(), svcCtx)
		resp, err := l.GetUserById(&req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
