package cert

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

func (l *ListLogic) List(req *types.CertListReq) (resp *types.CertListResp, err error) {
	dbConn := withWhere(req)
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	groups, count, err := l.svcCtx.CertsModel.List(l.ctx, dbConn, req.Page, req.Size)
	if err != nil {
		logx.Error(err)
		err = errors.DbError()
		return nil, err
	}
	resp = &types.CertListResp{
		Total: count,
		List:  make([]*types.CertItem, 0, len(groups)),
	}

	for _, v := range groups {
		itme := &types.CertItem{
			Id:      v.Id,
			Name:    v.Name,
			Desc:    v.Desc,
			Account: v.Account,
			Type:    uint8(v.Type),
		}
		dbConn := squirrel.Select().Where(squirrel.Eq{"groupId": v.Id})
		itme.HostNum, err = l.svcCtx.RulesModel.Count(l.ctx, dbConn)
		if err != nil {
			logx.Error(err)
			err = errors.DbError()
			return nil, err
		}
		resp.List = append(resp.List, itme)
	}
	return
}

func withWhere(data *types.CertListReq) squirrel.SelectBuilder {
	dbConn := squirrel.Select()
	if data.Name != "" {
		dbConn = dbConn.Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", data.Name)})
	}
	return dbConn
}
