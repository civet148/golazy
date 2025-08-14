package user

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1/user"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 修改用户
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param EditUser body types.EditUserReq true "request params description"
// @Success 200 {object} types.EditUserRsp
// @Router /api/v1/user/edit [post]
func EditUserHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.EditUserReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := user.NewEditUserLogic(c, svcCtx)
		resp, err := l.EditUser(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
