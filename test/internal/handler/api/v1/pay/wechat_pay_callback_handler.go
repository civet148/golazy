package pay

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/types"

	"github.com/civet148/log"
	"test/internal/logic/api/v1/pay"
	"test/internal/svc"
)

// @Summary 微信支付回调
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param WechatPayCallback body types.WechatPayCallbackReq true "params description"
// @Success 200 {object} nil
// @Router /api/v1/pay/wechat/pay/:tid [post]
func WechatPayCallbackHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.WechatPayCallbackReq
		if err := c.ShouldBindUri(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := pay.NewWechatPayCallbackLogic(c, svcCtx)

		err := l.WechatPayCallback(c, &req)
		if err != nil {
			log.Errorf("call WechatPayCallback failed, err: %v", err.Error())
			return
		}

	}
}
