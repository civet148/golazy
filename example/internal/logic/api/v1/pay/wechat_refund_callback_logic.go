package pay

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type WechatRefundCallbackLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 微信退款回调
func NewWechatRefundCallbackLogic(c *gin.Context, svcCtx *svc.ServiceContext) *WechatRefundCallbackLogic {
	return &WechatRefundCallbackLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *WechatRefundCallbackLogic) WechatRefundCallback(ctx context.Context, req *types.WechatRefundCallbackReq) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
