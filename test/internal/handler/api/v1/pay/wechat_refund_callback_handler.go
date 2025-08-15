package pay

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/types"

	"github.com/civet148/log"
	"test/internal/logic/api/v1/pay"
	"test/internal/svc"
)

// @Summary 微信退款回调
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param WechatRefundCallback body types.WechatRefundCallbackReq true "params description"
// @Success 200 {object} nil
// @Router /api/v1/pay/wechat/refund/:tid [post]
func WechatRefundCallbackHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.WechatRefundCallbackReq
		if err := c.ShouldBindUri(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := pay.NewWechatRefundCallbackLogic(c, svcCtx)

		err := l.WechatRefundCallback(c, &req)
		if err != nil {
			log.Errorf("call WechatRefundCallback failed, err: %v", err.Error())
			return
		}

	}
}
