package pay

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1/pay"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 微信支付回调
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param WechatPayCallbackHandler body types.WechatPayCallbackReq true "request params description"
// @Success 200 {object} types.WechatPayCallbackRsp
// @Router /api/v1/pay/wechat/{id:[0-9]+} [get]
func WechatPayCallbackHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.WechatPayCallbackReq
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := pay.NewWechatPayCallbackLogic(c, svcCtx)
		resp, err := l.WechatPayCallback(c, &req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
