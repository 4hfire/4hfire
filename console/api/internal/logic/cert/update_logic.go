package cert

import (
	"4hfire/common/errors"
	"4hfire/model"
	"context"
	"unicode/utf8"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.CertUpdateReq) error {
	cert, err := l.svcCtx.CertsModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errors.DbError()
	}
	if utf8.RuneCountInString(req.Name) > 64 {
		return errors.ParamsError("凭证名称长度不合规，长度需要小于64位")
	}
	if utf8.RuneCountInString(req.Desc) > 255 {
		return errors.ParamsError("凭证备注长度不合规，长度需要小于255位")
	}
	if utf8.RuneCountInString(req.Account) > 64 {
		return errors.ParamsError("凭证账号长度不合规，长度需要小于255位")
	}
	if utf8.RuneCountInString(req.Password) > 64 {
		return errors.ParamsError("凭证密码长度不合规，长度需要小于255位")
	}
	if utf8.RuneCountInString(req.Secret) > 1024 {
		return errors.ParamsError("凭证私钥长度不合规，长度需要小于255位")
	}
	if len(req.Account) == 0 && ((req.Type == 1 && len(req.Password) == 0) ||
		(req.Type == 2 && len(req.Secret) == 0)) {
		return errors.ParamsError()
	}
	if cert.Name != req.Name {
		if _, err := l.svcCtx.CertsModel.FindOneByName(l.ctx, req.Name); err != model.ErrNotFound {
			return errors.ParamsError("该凭证名称已存在")
		}
	}
	cert.Name = req.Name
	cert.Desc = req.Desc
	cert.Type = uint64(req.Type)
	cert.Account = req.Account
	cert.Password = req.Password
	cert.Secret = req.Secret
	err = l.svcCtx.CertsModel.Update(l.ctx, cert)

	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
