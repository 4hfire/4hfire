package staff

import (
	"4hfire/common/errors"
	"4hfire/common/otp"
	"context"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenOTPInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理员获取otp开启信息
func NewOpenOTPInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenOTPInfoLogic {
	return &OpenOTPInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenOTPInfoLogic) OpenOTPInfo() (resp *types.OpenOTPInfo, err error) {
	account := l.ctx.Value("account").(string)
	if account == "" {
		return &types.OpenOTPInfo{}, errors.GetOTPInfoError()
	}

	u, err := l.svcCtx.AdminsModel.FindOneByAccount(l.ctx, account)
	if err != nil {
		l.Logger.Error(err)
		return &types.OpenOTPInfo{}, errors.DbError()
	}
	u.Code = otp.RandNewStr(32)
	err = l.svcCtx.AdminsModel.Update(l.ctx, u)
	if err != nil {
		return &types.OpenOTPInfo{}, errors.DbError()
	}

	url, _ := otp.Create(u.Account, u.Code)
	return &types.OpenOTPInfo{Data: url}, nil
}
