package tags

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

func (l *OptionLogic) Option() (resp *types.TagOptionResp, err error) {

	tags, err := l.svcCtx.TagsModel.All(l.ctx)
	resp = &types.TagOptionResp{
		Item: make([]*types.TagOption, 0, len(tags)),
	}

	for _, v := range tags {
		itme := &types.TagOption{
			Id:   v.Id,
			Name: v.Name,
		}
		resp.Item = append(resp.Item, itme)
	}
	return
}
