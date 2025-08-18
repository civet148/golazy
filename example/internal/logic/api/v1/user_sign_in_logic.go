package v1

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type UserSignInLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewUserSignInLogic(c *gin.Context, svcCtx *svc.ServiceContext) *UserSignInLogic {
	return &UserSignInLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *UserSignInLogic) UserSignIn(ctx context.Context, req *types.UserSignInReq) (resp *types.UserSignInRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.UserSignInRsp{}, nil
}
