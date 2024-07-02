package staff

import (
	"4hfire/common/errors"
	"4hfire/model"
	"context"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理员删除管理账户
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.AdminDeleteReq) error {

	if req.Uid == l.ctx.Value("uid").(string) {
		return errors.CommitSuicide()
	}
	user, err := l.svcCtx.AdminsModel.FindOneByUid(l.ctx, req.Uid)
	if err != nil && err != model.ErrNotFound {
		err = errors.DbError()
		return err
	}
	if err == model.ErrNotFound {
		err = errors.ErrorNotFound()
		return err
	}

	err = l.svcCtx.AdminsModel.Delete(l.ctx, user.Id)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	//	用户下线
	_ = kickOutUser(l.ctx, l.svcCtx, req.Uid)
	return nil
}

func kickOutUser(ctx context.Context, svcCtx *svc.ServiceContext, uid string) (err error) {
	//	获取上个token
	formerToken, err := svcCtx.JwtInstance.Load(ctx, uid)
	if err != nil {
		logx.Error("获取上个token失败 ", err)
	}
	//	上个token存在 则废弃掉
	if formerToken != "" {
		err = svcCtx.JwtInstance.DiscardWithDelete(formerToken)
		//	废弃失败 返回错误
		if err != nil {
			logx.Error("废弃上个token失败 ", err)
		}
	}
	return
}
