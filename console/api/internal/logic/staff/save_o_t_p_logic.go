package staff

import (
	"4hfire/common/errors"
	"4hfire/common/otp"
	"4hfire/model"
	"context"
	"strconv"
	"time"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOTPLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 保存otp信息
func NewSaveOTPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOTPLogic {
	return &SaveOTPLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOTPLogic) SaveOTP(req *types.SaveOTPReq) error {

	account := l.ctx.Value("account").(string)
	if account == "" {
		return errors.SystemError()
	}

	u, err := l.svcCtx.AdminsModel.FindOneByAccount(l.ctx, account)
	if err != nil {
		l.Logger.Error(err)
		return errors.DbError()
	}

	if u.Code == "" {
		return errors.SystemError()
	}

	code, err := strconv.Atoi(req.Code)
	if err != nil {
		return errors.ParamsError()
	}
	otpCode := otp.ReturnCode(u.Code)
	if uint32(code) != otpCode {
		return errors.OTPCodeError()
	}

	u.Otp = model.AdminEnableOTP
	u.UpdatedAt = time.Now().Unix()

	if err := l.svcCtx.AdminsModel.Update(l.ctx, u); err != nil {
		l.Logger.Error(err)
		return errors.DbError()
	}

	return nil
}
