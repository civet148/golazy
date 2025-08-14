package user

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1/user"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 添加用户
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param AddUser body types.AddUserReq true "request params description"
// @Success 200 {object} types.AddUserRsp
// @Router /api/v1/user/add [put]
func AddUserHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.AddUserReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := user.NewAddUserLogic(c, svcCtx)
		resp, err := l.AddUser(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
