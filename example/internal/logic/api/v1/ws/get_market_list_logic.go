package ws

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
)

type GetMarketListLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 市场行情websocket
func NewGetMarketListLogic(c *gin.Context, svcCtx *svc.ServiceContext) *GetMarketListLogic {
	return &GetMarketListLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *GetMarketListLogic) GetMarketList(ctx context.Context) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
