package staff

import (
	"4hfire/common/errors"
	"4hfire/model"
	"context"
	"time"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseOTPLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理员关闭otp
func NewCloseOTPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseOTPLogic {
	return &CloseOTPLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CloseOTPLogic) CloseOTP(req *types.CloseOTPReq) error {
	u, err := l.svcCtx.AdminsModel.FindOneByUid(l.ctx, req.UID)
	if err != nil {
		l.Logger.Error(err)
		return errors.DbError()
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
