package host

import (
	"4hfire/common/errors"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.HostDeleteReq) error {
	err := l.svcCtx.MysqlConn.TransactCtx(l.ctx, func(ctx context.Context, s sqlx.Session) error {
		err := l.svcCtx.HostsModel.DeleteSession(l.ctx, req.Id, s)
		if err != nil {
			logx.Error(err)
			return errors.SystemError()
		}
		err = l.svcCtx.HostCertsModel.DeleteSession(l.ctx, req.Id, s)
		if err != nil {
			logx.Error(err)
			return errors.SystemError()
		}
		return nil
	})
	if err != nil {
		logx.Error(err)
		return errors.SystemError()
	}
	return nil
}
