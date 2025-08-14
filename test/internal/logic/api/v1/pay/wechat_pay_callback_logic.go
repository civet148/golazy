package pay

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
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

func (l *WechatPayCallbackLogic) WechatPayCallback(ctx context.Context, req *types.WechatPayCallbackReq) (resp *types.WechatPayCallbackRsp, err error) {
	// todo: add your logic here and delete this line
	// you can call ctx.(*gin.Context) convert to gin context
	return &types.WechatPayCallbackRsp{}, nil
}
