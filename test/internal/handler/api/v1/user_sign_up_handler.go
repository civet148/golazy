package v1

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 用户注册
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param UserSignUp body types.UserSignUpReq true "request params description"
// @Success 200 {object} types.UserSignUpRsp
// @Router /api/v1/sign_up [post]
func UserSignUpHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignUpReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := v1.NewUserSignUpLogic(c, svcCtx)
		resp, err := l.UserSignUp(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
