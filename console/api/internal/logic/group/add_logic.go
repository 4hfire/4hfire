package group

import (
	"4hfire/common/errors"
	"4hfire/common/utils"
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

func (l *AddLogic) Add(req *types.RuleGroupAddReq) error {
	if _, err := l.svcCtx.RuleGroupsModel.FindOneByName(l.ctx, req.Name); err != model.ErrNotFound {
		logx.Error(err)
		err = errors.AccountRepeatedError()
		return err
	}
	if utf8.RuneCountInString(req.Name) > 64 {
		return errors.ParamsError("规则分组名称长度不合规，长度需要小于64位")
	}
	if utf8.RuneCountInString(req.Desc) > 255 {
		return errors.ParamsError("规则分组名称长度不合规，长度需要小于255位")
	}
	data := &model.RuleGroups{
		Name: req.Name,
		Desc: req.Desc,
		Tags: utils.ToJSONString(req.Tags),
	}
	_, err := l.svcCtx.RuleGroupsModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
