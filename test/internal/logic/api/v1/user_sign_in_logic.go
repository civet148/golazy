package v1

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
)

type UserSignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewUserSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignInLogic {
	return &UserSignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignInLogic) UserSignIn(ctx context.Context, req *types.UserSignInReq) (resp *types.UserSignInRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.UserSignInRsp{}, nil
}
