package group

import (
	"4hfire/common/errors"
	"context"
	"github.com/Masterminds/squirrel"

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

func (l *DeleteLogic) Delete(req *types.RuleGroupDeleteReq) error {

	dbConn := squirrel.Select().Where(squirrel.Eq{"groupId": req.Id})
	c, err := l.svcCtx.RulesModel.Count(l.ctx, dbConn)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	if c > 0 {
		return errors.DiyError("存在关联的规则，请解绑后删除")
	}
	err = l.svcCtx.RuleGroupsModel.Delete(l.ctx, req.Id)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
