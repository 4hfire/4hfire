package auth

import (
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"4hfire/common/errors"
	"4hfire/model"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理-管理用户详情
func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info() (resp *types.AdminInfoResp, err error) {
	uid := fmt.Sprintf("%v", l.ctx.Value("uid"))
	user, err := l.svcCtx.AdminsModel.FindOneByUid(l.ctx, uid)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errors.DbError()
	}
	if err == model.ErrNotFound {
		return nil, errors.ErrorNotFound()
	}
	avatar, err := l.svcCtx.AdminAvatarsModel.FindOne(l.ctx, int64(user.Id))
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errors.DbError()
	}
	if err == model.ErrNotFound {
		return nil, errors.ErrorNotFound()
	}
	resp = &types.AdminInfoResp{
		Uid:        user.Uid,
		Phone:      user.Phone,
		Account:    user.Account,
		Avatar:     avatar.Avatar.String,
		CreateTime: user.CreatedAt,
		UpdateTime: user.UpdatedAt,
		Email:      user.Email,
		Name:       user.Name,
	}

	return
}
