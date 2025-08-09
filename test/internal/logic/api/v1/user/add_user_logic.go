package user

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加用户
func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.AddUserRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
