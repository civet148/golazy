package ws

import (
	"context"

	"github.com/gin-gonic/gin"

	"test/internal/svc"
)

type WsMarketListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 市场行情（websocket方式）
func NewWsMarketListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsMarketListLogic {
	return &WsMarketListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsMarketListLogic) WsMarketList(c *gin.Context) error {
	// todo: add your logic here and delete this line

	return nil
}
