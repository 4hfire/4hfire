package staff

import (
	"4hfire/common/errors"
	"4hfire/common/utils"
	"4hfire/model"
	"context"
	"time"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 控制台管理员新增
func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AdminAdduserReq) error {
	if _, err := l.svcCtx.AdminsModel.FindOneByAccount(l.ctx, req.Account); err != model.ErrNotFound {
		logx.Error(err)
		err = errors.AccountRepeatedError()
		return err
	}

	if req.Phone != "" && !utils.VerifyPhone(req.Phone) {
		return errors.ValidatePhoneError()
	}
	if len(req.Account) < 6 || len(req.Account) > 26 {
		return errors.ParamsError("用户名长度不合规，长度需要大于6并且小于26位")
	}
	if len(req.Password) != 32 {
		return errors.ParamsError()
	}

	t := time.Now().Unix()
	data := &model.Admins{
		Phone:     req.Phone,
		Account:   req.Account,
		Password:  utils.MD5(req.Password),
		CreatedAt: t,
		UpdatedAt: t,
		IsDisable: model.AdminIsNotDisabled,
		Email:     req.Email,
		Uid:       uuid.NewV4().String(),
		Name:      req.Name,
		Otp:       model.AdminEnableOTP,
	}
	_, err := l.svcCtx.AdminsModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
