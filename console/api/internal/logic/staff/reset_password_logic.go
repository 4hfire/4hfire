package staff

import (
	"4hfire/common/errors"
	"4hfire/common/utils"
	"4hfire/model"
	"context"
	"time"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理员修改密码
func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.AdminResetPasswordReq) error {

	//	前后密码不一致 长度不合规
	if req.NewPassword != req.ReNewPassword || len(req.NewPassword) != 32 {
		return errors.ParamsError()
	}
	user, err := l.svcCtx.AdminsModel.FindOneByUid(l.ctx, req.Uid)
	if err != nil && err != model.ErrNotFound {
		return errors.DbError()
	}
	if err == model.ErrNotFound {
		return errors.ErrorNotFound()
	}

	user.Password = utils.MD5(req.NewPassword)
	user.UpdatedAt = time.Now().Unix()

	err = l.svcCtx.AdminsModel.Update(l.ctx, user)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}

	//	用户下线
	_ = kickOutUser(l.ctx, l.svcCtx, req.Uid)
	return nil
}
