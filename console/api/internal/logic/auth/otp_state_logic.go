package auth

import (
	"4hfire/model"
	"context"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OtpStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录时查看otp开启状态
func NewOtpStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtpStateLogic {
	return &OtpStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OtpStateLogic) OtpState(req *types.OTPStateReq) (resp *types.OTPStateResp, err error) {
	u, err := l.svcCtx.AdminsModel.FindOneByAccount(l.ctx, req.Account)
	if err != nil {
		l.Logger.Error(err)
		return &types.OTPStateResp{Enable: false}, nil
	}
	return &types.OTPStateResp{Enable: u.Otp == model.AdminEnableOTP}, nil
}
