package user

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type GetUserListLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 用户列表
func NewGetUserListLogic(c *gin.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(ctx context.Context, req *types.GetUserListReq) (resp *types.GetUserListRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.GetUserListRsp{}, nil
}
