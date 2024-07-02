package rule

import (
	"4hfire/common/errors"
	"4hfire/model"
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

func (l *ListLogic) List(req *types.RuleListReq) (resp *types.RuleListResp, err error) {
	dbConn := withWhere(req)
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	tags, count, err := l.svcCtx.RulesModel.List(l.ctx, dbConn, req.Page, req.Size)
	if err != nil {
		logx.Error(err)
		err = errors.DbError()
		return nil, err
	}
	resp = &types.RuleListResp{
		Total: count,
		List:  make([]*types.RuleItem, 0, len(tags)),
	}
	gnm := map[uint64]string{}
	for _, v := range tags {
		itme := &types.RuleItem{
			Id:        v.Id,
			Name:      v.Name,
			Interface: v.Interface,
			GroupId:   v.GroupId,
			Enable:    v.Enable == model.RuleEnable,
			SIP:       v.Sip,
			SPort:     v.Sport,
			DIP:       v.Dip,
			DPort:     v.Dport,
			Option:    "accept",
		}
		if v.Option == model.RuleOptionDrop {
			itme.Option = "drop"
		}
		if g, ok := gnm[v.GroupId]; !ok {
			group, err := l.svcCtx.RuleGroupsModel.FindOne(l.ctx, v.GroupId)
			if err != nil {
				logx.Error(err)
				err = errors.DbError()
				return nil, err
			}
			gnm[v.GroupId] = group.Name
			itme.GroupName = group.Name
		} else {
			itme.GroupName = g
		}
		resp.List = append(resp.List, itme)
	}
	return
}

func withWhere(data *types.RuleListReq) squirrel.SelectBuilder {
	dbConn := squirrel.Select()
	if data.Name != "" {
		dbConn = dbConn.Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", data.Name)})
	}
	if data.GroupId != 0 {
		dbConn = dbConn.Where(squirrel.Eq{"groupId": data.GroupId})
	}
	return dbConn
}
