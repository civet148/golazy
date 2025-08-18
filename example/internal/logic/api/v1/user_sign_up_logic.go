package v1

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type UserSignUpLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewUserSignUpLogic(c *gin.Context, svcCtx *svc.ServiceContext) *UserSignUpLogic {
	return &UserSignUpLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *UserSignUpLogic) UserSignUp(ctx context.Context, req *types.UserSignUpReq) (resp *types.UserSignUpRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.UserSignUpRsp{}, nil
}
