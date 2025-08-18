package pay

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type WechatPayCallbackLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 微信支付回调
func NewWechatPayCallbackLogic(c *gin.Context, svcCtx *svc.ServiceContext) *WechatPayCallbackLogic {
	return &WechatPayCallbackLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *WechatPayCallbackLogic) WechatPayCallback(ctx context.Context, req *types.WechatPayCallbackReq) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
