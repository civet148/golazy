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

func (l *EditUserLogic) EditUser(req *types.EditUserReq) (resp *types.EditUserRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
