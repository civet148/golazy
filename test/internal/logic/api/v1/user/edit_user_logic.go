package user

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
)

type EditUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户
func NewEditUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserLogic {
	return &EditUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserLogic) EditUser(ctx context.Context, req *types.EditUserReq) (resp *types.EditUserRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.EditUserRsp{}, nil
}
