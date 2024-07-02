package tags

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

func (l *UpdateLogic) Update(req *types.TagUpdateReq) error {
	tag, err := l.svcCtx.TagsModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errors.DbError()
	}

	if utf8.RuneCountInString(req.Name) > 64 {
		return errors.ParamsError("标签名称长度不合规，长度需要小于64位")

	}
	if utf8.RuneCountInString(req.Desc) > 255 {
		return errors.ParamsError("标签名称长度不合规，长度需要小于255位")
	}
	if tag.Name != req.Name {
		if _, err := l.svcCtx.TagsModel.FindOneByName(l.ctx, req.Name); err != model.ErrNotFound {
			return errors.ParamsError("该标签名称已存在")
		}
	}
	tag.Name = req.Name
	tag.Desc = req.Desc
	err = l.svcCtx.TagsModel.Update(l.ctx, tag)

	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
