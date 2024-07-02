package auth

import (
	"context"

	"4hfire/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理用户退出
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() error {
	formerToken, err := l.svcCtx.JwtInstance.Load(l.ctx, l.ctx.Value("uid").(string))
	if err != nil {
		logx.Error(err)
		return nil
	}
	//	上个token存在 则废弃掉
	if formerToken != "" {
		err = l.svcCtx.JwtInstance.DiscardWithDelete(formerToken)
		//	废弃失败 返回错误
		if err != nil {
			logx.Error(err)
			return nil
		}
	}

	return nil
}
