package pay

import (
	"context"

	"example/internal/svc"
	"example/internal/types"
)

type WechatPayCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 微信支付回调
func NewWechatPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatPayCallbackLogic {
	return &WechatPayCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatPayCallbackLogic) WechatPayCallback(ctx context.Context, req *types.WechatPayCallbackReq) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
