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
// @Param DeleteUser body types.DeleteUserReq true "params description"
// @Success 200 {object} types.DeleteUserRsp
// @Router /api/v1/user/delete [delete]
func DeleteUserHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.DeleteUserReq
		if err := c.ShouldBind(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := user.NewDeleteUserLogic(c, svcCtx)

		resp, err := l.DeleteUser(c, &req)
		if err != nil {
			log.Errorf("call DeleteUser failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
