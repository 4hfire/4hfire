package group

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

func (l *UpdateLogic) Update(req *types.RuleGroupUpdateReq) error {
	group, err := l.svcCtx.RuleGroupsModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errors.DbError()
	}

	if utf8.RuneCountInString(req.Name) > 64 {
		return errors.ParamsError("规则分组名称长度不合规，长度需要小于64位")

	}
	if utf8.RuneCountInString(req.Desc) > 255 {
		return errors.ParamsError("规则分组名称长度不合规，长度需要小于255位")
	}
	if group.Name != req.Name {
		if _, err := l.svcCtx.RuleGroupsModel.FindOneByName(l.ctx, req.Name); err != model.ErrNotFound {
			return errors.ParamsError("该规则分组名称已存在")
		}
	}
	group.Name = req.Name
	group.Desc = req.Desc
	err = l.svcCtx.RuleGroupsModel.Update(l.ctx, group)

	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
