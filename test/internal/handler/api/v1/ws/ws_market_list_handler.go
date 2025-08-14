package ws

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"test/internal/logic/api/v1/ws"
	"test/internal/svc"
)

// @Summary 市场行情websocket
// @Description
// @Tags
// @Accept plain
// @Produce plain
// @Param WsMarketList body string true "params description"
// @Success 200 {string} nil
// @Router /api/v1/ws/market [get]
func WsMarketListHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		l := ws.NewWsMarketListLogic(c, svcCtx)
		err := l.WsMarketList(c)
		if err != nil {
			log.Errorf("call WsMarketList failed, err: %v", err.Error())
		}
		c.Abort()

	}
}
