package user

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据ID查询用户
func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIdReq) (resp *types.GetUserByIdRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
