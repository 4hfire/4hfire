package staff

import (
	"4hfire/common/errors"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 管理员通过账号获取管理员列表
func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.AdminListReq) (resp *types.AdminListRes, err error) {
	dbConn := withAdminUserWhere(req)
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	users, count, err := l.svcCtx.AdminsModel.List(l.ctx, dbConn, req.Page, req.Size)
	if err != nil {
		logx.Error(err)
		err = errors.DbError()
		return nil, err
	}
	resp = &types.AdminListRes{
		Total: count,
		List:  make([]*types.AdminItem, 0, len(users)),
	}

	for _, v := range users {
		itme := &types.AdminItem{
			Uid:        v.Uid,
			Phone:      v.Phone,
			Account:    v.Account,
			CreateTime: v.CreatedAt,
			IsDisable:  int64(v.IsDisable),
			UpdateTime: v.UpdatedAt,
			Email:      v.Email,
			Name:       v.Name,
		}
		resp.List = append(resp.List, itme)
	}
	return
}

func withAdminUserWhere(data *types.AdminListReq) squirrel.SelectBuilder {
	dbConn := squirrel.Select()
	// .Where(sq.Eq{"uid": uid})
	if data.Account != "" {
		dbConn = dbConn.Where(squirrel.Like{"account": fmt.Sprintf("%%%s%%", data.Account)})
	}
	if data.Name != "" {
		dbConn = dbConn.Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", data.Name)})
	}
	if data.Phone != "" {
		dbConn = dbConn.Where(squirrel.Like{"phone": fmt.Sprintf("%%%s%%", data.Phone)})
	}

	return dbConn
}
