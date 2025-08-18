package user

import (
	"context"

	"github.com/gin-gonic/gin"

	"example/internal/svc"
	"example/internal/types"
)

type EditUserLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 修改用户
func NewEditUserLogic(c *gin.Context, svcCtx *svc.ServiceContext) *EditUserLogic {
	return &EditUserLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *EditUserLogic) EditUser(ctx context.Context, req *types.EditUserReq) (resp *types.EditUserRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.EditUserRsp{}, nil
}
