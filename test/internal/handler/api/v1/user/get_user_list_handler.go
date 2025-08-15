package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/types"

	"github.com/civet148/log"
	"test/internal/logic/api/v1/user"
	"test/internal/svc"
)

// @Summary 用户列表
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param GetUserList body types.GetUserListReq true "params description"
// @Success 200 {object} types.GetUserListRsp
// @Router /api/v1/user/list [get]
func GetUserListHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.GetUserListReq
		if err := c.ShouldBind(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := user.NewGetUserListLogic(c, svcCtx)

		resp, err := l.GetUserList(c, &req)
		if err != nil {
			log.Errorf("call GetUserList failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
