package user

import (
	"context"

	"github.com/pz2147/p-gateway-1/internal/svc"
	"github.com/pz2147/p-gateway-1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.EmptyReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
