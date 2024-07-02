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

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.CertAddReq) error {
	if _, err := l.svcCtx.CertsModel.FindOneByName(l.ctx, req.Name); err != model.ErrNotFound {
		logx.Error(err)
		err = errors.CertAddRepatedError()
		return err
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

	data := &model.Certs{
		Name:     req.Name,
		Desc:     req.Desc,
		Account:  req.Account,
		Password: req.Password,
		Secret:   req.Secret,
		Type:     uint64(req.Type),
	}
	_, err := l.svcCtx.CertsModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
