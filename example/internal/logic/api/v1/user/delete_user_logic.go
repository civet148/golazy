package user

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type DeleteUserLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewDeleteUserLogic(c *gin.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(ctx context.Context, req *types.DeleteUserReq) (resp *types.DeleteUserRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.DeleteUserRsp{}, nil
}
