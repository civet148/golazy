package ws

import (
	"context"
	"github.com/gin-gonic/gin"
	"test/internal/logic/api/v1/ws"
	"test/internal/svc"
)

// 市场行情（websocket方式）
func WsMarketListHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		l := ws.NewWsMarketListLogic(context.Background(), svcCtx)
		_ = l.WsMarketList(c)

	}
}
