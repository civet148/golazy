package v1

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
)

type UserSignOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户退出登录
func NewUserSignOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignOutLogic {
	return &UserSignOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignOutLogic) UserSignOut(req *types.UserSignOutReq) (resp *types.UserSignOutRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
