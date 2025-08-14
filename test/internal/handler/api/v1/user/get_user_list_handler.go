package user

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1/user"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 用户列表
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param GetUserList body types.GetUserListReq true "request params description"
// @Success 200 {object} types.GetUserListRsp
// @Router /api/v1/user/list [get]
func GetUserListHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.GetUserListReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := user.NewGetUserListLogic(c, svcCtx)
		resp, err := l.GetUserList(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
