package host

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

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.HostListReq) (resp *types.HostListResp, err error) {
	dbConn := withWhere(req)
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	groups, count, err := l.svcCtx.HostsModel.List(l.ctx, dbConn, req.Page, req.Size)
	if err != nil {
		logx.Error(err)
		err = errors.DbError()
		return nil, err
	}
	resp = &types.HostListResp{
		Total: count,
		List:  make([]*types.HostItem, 0, len(groups)),
	}

	for _, v := range groups {
		itme := &types.HostItem{
			Id:           v.Id,
			Name:         v.Name,
			Desc:         v.Desc,
			Address:      v.Address,
			CPU:          v.Cpu,
			Mem:          v.Mem,
			LastUploadAt: v.UpdatedAt,
			State:        v.State,
		}
		resp.List = append(resp.List, itme)
	}
	return
}

func withWhere(data *types.HostListReq) squirrel.SelectBuilder {
	dbConn := squirrel.Select()
	if data.Name != "" {
		dbConn = dbConn.Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", data.Name)})
	}
	if data.Address != "" {
		dbConn = dbConn.Where(squirrel.Like{"address": fmt.Sprintf("%%%s%%", data.Address)})
	}
	return dbConn
}
