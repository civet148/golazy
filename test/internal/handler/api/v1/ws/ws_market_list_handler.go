package ws

import (
	"context"
	"github.com/gin-gonic/gin"
	"test/internal/logic/api/v1/ws"
	"test/internal/svc"
)

// @Summary 市场行情（websocket方式）
// @Description
// @Tags
// @Accept plain
// @Produce plain
// @Param WsMarketListHandler body string true "request params description"
// @Success 200 {string} string
// @Router /api/v1/ws/market [get]
func WsMarketListHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		l := ws.NewWsMarketListLogic(context.Background(), svcCtx)
		_ = l.WsMarketList(c)

	}
}
