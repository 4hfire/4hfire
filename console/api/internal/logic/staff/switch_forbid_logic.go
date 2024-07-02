package staff

import (
	"4hfire/common/errors"
	"4hfire/model"
	"context"
	"time"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchForbidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理员账户切换账号禁用状态
func NewSwitchForbidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchForbidLogic {
	return &SwitchForbidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwitchForbidLogic) SwitchForbid(req *types.AdminForbidReq) error {
	if req.Uid == l.ctx.Value("uid").(string) {
		return errors.ExistCanNotOperate()
		// return errors.NewErrCodeMsg(errors.INNER_ERROR, "无法删除自己账户")
	}
	user, err := l.svcCtx.AdminsModel.FindOneByUid(l.ctx, req.Uid)
	if err != nil && err != model.ErrNotFound {
		return errors.DbError()
	}
	if err == model.ErrNotFound {
		return errors.ErrorNotFound()
	}

	if user.IsDisable == model.AdminIsDisabled {
		user.IsDisable = model.AdminIsNotDisabled
	} else {
		user.IsDisable = model.AdminIsDisabled
		//	用户下线
		_ = kickOutUser(l.ctx, l.svcCtx, req.Uid)
	}
	user.UpdatedAt = time.Now().Unix()
	err = l.svcCtx.AdminsModel.Update(l.ctx, user)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
