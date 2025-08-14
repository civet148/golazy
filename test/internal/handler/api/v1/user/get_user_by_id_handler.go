package user

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1/user"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 根据ID查询用户
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param GetUserById body types.GetUserByIdReq true "request params description"
// @Success 200 {object} types.GetUserByIdRsp
// @Router /api/v1/user/:id [get]
func GetUserByIdHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.GetUserByIdReq
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := user.NewGetUserByIdLogic(c, svcCtx)
		resp, err := l.GetUserById(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
