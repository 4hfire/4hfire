package auth

import (
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"4hfire/common/errors"
	"4hfire/common/lib/jwt"
	"4hfire/common/otp"
	"4hfire/common/utils"
	"4hfire/model"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 鉴权登录-管理登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	admin, err := l.svcCtx.AdminsModel.FindOneByAccount(l.ctx, req.Username)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errors.DbError()
	}
	if err == model.ErrNotFound {
		if req.Username == model.DEFAULT_ADMIN_USERNAME {
			if err = l.svcCtx.AdminsModel.CreateDefaultAdmin(l.ctx); err != nil {
				logx.Error(err)
				return nil, errors.DbError()
			}
		}
		if req.Password != utils.MD5(model.DEFAULT_ADMIN_PASSWORD) {
			return nil, errors.LoginError()
		}
		return nil, errors.LoginError()
	}
	if admin.Password != utils.MD5(req.Password) {
		return nil, errors.LoginError()
	}
	if admin.IsDisable == model.AdminIsDisabled {
		return nil, errors.ForbidError()
	}
	if admin.Otp != model.AdminEnableOTP {
		if req.OTPCode == "" || admin.Code == "" {
			err = errors.OTPCodeError()
			return nil, err
		}
		if !otp.Check(admin.Code, req.OTPCode) {
			err = errors.OTPCodeError()
			return nil, err
		}
	}
	admin.LastLoginTime = time.Now().Unix()
	_ = l.svcCtx.AdminsModel.Update(l.ctx, admin)
	formerToken, err := l.svcCtx.JwtInstance.Load(l.ctx, admin.Uid)

	if err != nil {
		logx.Error(err)
		return nil, errors.SystemError()
	}
	//	上个token存在 则废弃掉
	if formerToken != "" {
		err = l.svcCtx.JwtInstance.DiscardWithDelete(formerToken)
		//	废弃失败 返回错误
		if err != nil {
			logx.Error(err)
			return nil, errors.SystemError()
		}
	}

	//	生成新的token
	payload := make(jwt.Payload)
	payload["uid"] = admin.Uid
	payload["account"] = admin.Account
	if err != nil {
		logx.Error(err)
		return nil, errors.SystemError()
	}

	token, err := l.svcCtx.JwtInstance.Token(payload)
	//	签发新token失败
	if err != nil {
		logx.Error(err)
		return nil, errors.SystemError()
	}

	//	存储新的token
	err = l.svcCtx.JwtInstance.Store(l.ctx, admin.Uid, token)
	if err != nil {
		logx.Error(err)
		return nil, errors.SystemError()
	}
	return &types.LoginResp{
		ID:            admin.Id,
		Name:          admin.Name,
		Username:      admin.Account,
		LastLoginTime: time.Unix(admin.LastLoginTime, 0).Format("2006-01-02 15:04:05"),
		Token:         token,
	}, nil
}
