package rule

import (
	"4hfire/common/errors"
	"context"

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

func (l *DeleteLogic) Delete(req *types.RuleDeleteReq) error {
	err := l.svcCtx.RulesModel.Delete(l.ctx, req.Id)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
