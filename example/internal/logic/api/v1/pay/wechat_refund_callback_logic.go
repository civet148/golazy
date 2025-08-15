package pay

import (
	"context"

	"example/internal/svc"
	"example/internal/types"
)

type WechatRefundCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 微信退款回调
func NewWechatRefundCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatRefundCallbackLogic {
	return &WechatRefundCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatRefundCallbackLogic) WechatRefundCallback(ctx context.Context, req *types.WechatRefundCallbackReq) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
