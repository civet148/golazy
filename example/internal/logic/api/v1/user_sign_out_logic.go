package v1

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type UserSignOutLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 用户退出登录
func NewUserSignOutLogic(c *gin.Context, svcCtx *svc.ServiceContext) *UserSignOutLogic {
	return &UserSignOutLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *UserSignOutLogic) UserSignOut(ctx context.Context, req *types.UserSignOutReq) (resp *types.UserSignOutRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.UserSignOutRsp{}, nil
}
