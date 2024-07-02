package staff

import (
	"4hfire/common/errors"
	"4hfire/common/utils"
	"context"
	"time"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理员更新数据
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.AdminUpdateReq) error {
	user, err := l.svcCtx.AdminsModel.FindOneByUid(l.ctx, req.Uid)
	if err != nil {
		return errors.DbError()
	}

	if req.Phone != "" && !utils.VerifyPhone(req.Phone) {
		err = errors.ValidatePhoneError()
		return err
	}

	user.Email = req.Email
	user.Phone = req.Phone

	//	更改了用户的密码
	if req.Password != "" {
		user.Password = utils.MD5(req.Password)
		//	用户下线
		_ = kickOutUser(l.ctx, l.svcCtx, req.Uid)
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	user.UpdatedAt = time.Now().Unix()
	err = l.svcCtx.AdminsModel.Update(l.ctx, user)

	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
