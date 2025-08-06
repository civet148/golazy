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

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.GetUserListRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
