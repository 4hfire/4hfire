package tags

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

func (l *ListLogic) List(req *types.TagListReq) (resp *types.TagListResp, err error) {
	dbConn := withWhere(req)
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	tags, count, err := l.svcCtx.TagsModel.List(l.ctx, dbConn, req.Page, req.Size)
	if err != nil {
		logx.Error(err)
		err = errors.DbError()
		return nil, err
	}
	resp = &types.TagListResp{
		Total: count,
		List:  make([]*types.TagItem, 0, len(tags)),
	}

	for _, v := range tags {
		itme := &types.TagItem{
			Id:   v.Id,
			Name: v.Name,
			Desc: v.Desc,
			//todo 查询关联数量
		}
		dbConn := squirrel.Select().Where(fmt.Sprintf("JSON_CONTAINS(tags, '%d', '$');", v.Id))
		itme.RuleGroupNum, err = l.svcCtx.RuleGroupsModel.Count(l.ctx, dbConn)
		if err != nil {
			logx.Error(err)
			err = errors.DbError()
			return nil, err
		}
		itme.RuleNum, err = l.svcCtx.RulesModel.Count(l.ctx, dbConn)
		if err != nil {
			logx.Error(err)
			err = errors.DbError()
			return nil, err
		}
		resp.List = append(resp.List, itme)
	}
	return
}

func withWhere(data *types.TagListReq) squirrel.SelectBuilder {
	dbConn := squirrel.Select()
	if data.Name != "" {
		dbConn = dbConn.Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", data.Name)})
	}
	return dbConn
}
