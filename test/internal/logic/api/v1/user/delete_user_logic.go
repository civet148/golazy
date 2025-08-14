package user

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(ctx context.Context, req *types.DeleteUserReq) (resp *types.DeleteUserRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.DeleteUserRsp{}, nil
}
