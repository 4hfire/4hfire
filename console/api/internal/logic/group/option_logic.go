package group

import (
	"context"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OptionLogic {
	return &OptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OptionLogic) Option() (resp *types.RuleGroupOptionResp, err error) {

	gs, err := l.svcCtx.RuleGroupsModel.All(l.ctx)
	resp = &types.RuleGroupOptionResp{
		Item: make([]*types.RuleGroupOption, 0, len(gs)),
	}

	for _, v := range gs {
		itme := &types.RuleGroupOption{
			Id:   v.Id,
			Name: v.Name,
		}
		resp.Item = append(resp.Item, itme)
	}
	return
}
