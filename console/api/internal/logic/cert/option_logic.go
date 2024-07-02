package cert

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

func (l *OptionLogic) Option() (resp *types.CertOptionResp, err error) {

	cs, err := l.svcCtx.CertsModel.All(l.ctx)
	resp = &types.CertOptionResp{
		Item: make([]*types.CertOption, 0, len(cs)),
	}

	for _, v := range cs {
		itme := &types.CertOption{
			Id:   v.Id,
			Name: v.Name,
		}
		resp.Item = append(resp.Item, itme)
	}
	return
}
