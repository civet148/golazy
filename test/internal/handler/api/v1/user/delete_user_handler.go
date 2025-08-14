package user

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1/user"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 删除用户
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param DeleteUser body types.DeleteUserReq true "request params description"
// @Success 200 {object} types.DeleteUserRsp
// @Router /api/v1/user/delete [delete]
func DeleteUserHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.DeleteUserReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := user.NewDeleteUserLogic(c, svcCtx)
		resp, err := l.DeleteUser(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
