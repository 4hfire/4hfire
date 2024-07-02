package host

import (
	"4hfire/common/errors"
	"4hfire/common/utils"
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

func (l *UpdateLogic) Update(req *types.HostUpdateReq) error {
	host, err := l.svcCtx.HostsModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errors.DbError()
	}
	if utf8.RuneCountInString(req.Address) > 32 {
		return errors.ParamsError("主机地址长度不合规，长度需要小于64位")
	}
	if utf8.RuneCountInString(req.Name) > 32 {
		return errors.ParamsError("主机名称长度不合规，长度需要小于64位")
	}
	if utf8.RuneCountInString(req.Desc) > 255 {
		return errors.ParamsError("主机备注长度不合规，长度需要小于255位")
	}
	host.Name = req.Name
	host.Desc = req.Desc
	host.Tags = utils.ToJSONString(req.Tags)

	if host.Address != req.Address && host.Address != "" {
		// todo： 命令下发 更改地址
		host.Address = req.Address
	}
	err = l.svcCtx.HostsModel.Update(l.ctx, host)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
