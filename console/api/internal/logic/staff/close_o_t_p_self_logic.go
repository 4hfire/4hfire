package staff

import (
	"4hfire/common/errors"
	"4hfire/model"
	"context"
	"time"

	"4hfire/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CloseOTPSelfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 关闭自己otp
func NewCloseOTPSelfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseOTPSelfLogic {
	return &CloseOTPSelfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CloseOTPSelfLogic) CloseOTPSelf() error {
	account := l.ctx.Value("account").(string)
	if account == "" {
		return errors.SystemError()
	}
	u, err := l.svcCtx.AdminsModel.FindOneByAccount(l.ctx, account)
	if err != nil {
		l.Logger.Error(err)
		err = errors.DbError()
		return err
	}

	u.Otp = model.AdminDisableOTP
	u.Code = ""
	u.UpdatedAt = time.Now().Unix()

	if err := l.svcCtx.AdminsModel.Update(l.ctx, u); err != nil {
		l.Logger.Error(err)
		err = errors.DbError()
		return err
	}
	return nil
}
