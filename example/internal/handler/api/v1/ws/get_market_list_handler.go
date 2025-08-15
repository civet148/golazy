package ws

import (
	"example/internal/logic/api/v1/ws"
	"example/internal/svc"
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
)

// @Summary 市场行情websocket
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param GetMarketList body string true "params description"
// @Success 200 {object} nil
// @Router /api/v1/ws/market/list [get]
func GetMarketListHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		l := ws.NewGetMarketListLogic(c, svcCtx)

		err := l.GetMarketList(c)
		if err != nil {
			log.Errorf("call GetMarketList failed, err: %v", err.Error())
			return
		}

	}
}
