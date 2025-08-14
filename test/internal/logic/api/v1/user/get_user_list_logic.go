package user

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户列表
func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(ctx context.Context, req *types.GetUserListReq) (resp *types.GetUserListRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.GetUserListRsp{}, nil
}
