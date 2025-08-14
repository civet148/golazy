package ws

import (
	"context"

	"test/internal/svc"
)

type WsMarketListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 市场行情websocket
func NewWsMarketListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsMarketListLogic {
	return &WsMarketListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsMarketListLogic) WsMarketList(ctx context.Context) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
