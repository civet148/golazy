package user

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type GetUserByIdLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 根据ID查询用户
func NewGetUserByIdLogic(c *gin.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(ctx context.Context, req *types.GetUserByIdReq) (resp *types.GetUserByIdRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.GetUserByIdRsp{}, nil
}
