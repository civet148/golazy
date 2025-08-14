package v1

import (
	"context"

	"test/internal/svc"
	"test/internal/types"
)

type UserSignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewUserSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignUpLogic {
	return &UserSignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignUpLogic) UserSignUp(ctx context.Context, req *types.UserSignUpReq) (resp *types.UserSignUpRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.UserSignUpRsp{}, nil
}
