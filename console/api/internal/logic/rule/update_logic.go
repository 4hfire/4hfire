package rule

import (
	"4hfire/common/errors"
	"4hfire/common/utils"
	"4hfire/model"
	"context"
	"unicode/utf8"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.RuleUpdateReq) error {
	r, err := l.svcCtx.RulesModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Error(err)
		err = errors.UpdateError()
		return err
	}
	if r.Name != req.Name {
		if _, err := l.svcCtx.RulesModel.FindOneByName(l.ctx, req.Name); err != model.ErrNotFound {
			logx.Error(err)
			err = errors.RuleAddRepatedError()
			return err
		}
	}
	if utf8.RuneCountInString(req.Name) > 64 {
		return errors.ParamsError("规则名称长度不合规，长度需要小于64位")
	}
	if utf8.RuneCountInString(req.Desc) > 255 {
		return errors.ParamsError("规则备注长度不合规，长度需要小于255位")
	}
	if len(req.Interface) > 32 {
		return errors.ParamsError("网卡长度不合规，长度需要小于32位")
	}
	if len(req.SIP) > 32 {
		return errors.ParamsError("原IP长度不合规，长度需要小于32位")
	}
	if len(req.SPort) > 255 {
		return errors.ParamsError("源端口长度不合规，长度需要小于255位")
	}
	if len(req.DIP) > 32 {
		return errors.ParamsError("目的IP长度不合规，长度需要小于32位")
	}
	if len(req.DPort) > 255 {
		return errors.ParamsError("目的端口长度不合规，长度需要小于255位")
	}
	if !checkPort(req.SPort) {
		return errors.ParamsError("源端口参数格式错误")
	}
	if !checkPort(req.DPort) {
		return errors.ParamsError("目的端口参数格式错误")
	}
	data := &model.Rules{
		Name:      req.Name,
		Desc:      req.Desc,
		GroupId:   req.GroupId,
		Interface: req.Interface,
		Sip:       req.SIP,
		Sport:     req.SPort,
		Dip:       req.DIP,
		Dport:     req.DPort,
		Enable:    model.RuleDisabled,
		Tags:      utils.ToJSONString(req.Tags),
	}
	if req.Enable {
		data.Enable = model.RuleEnable
	}
	switch req.Option {
	case "accept", "ACCEPT":
		data.Option = model.RuleOptionAccept
	case "drop", "DROP":
		data.Option = model.RuleOptionDrop
	default:
		return errors.ParamsError()
	}
	err = l.svcCtx.RulesModel.Update(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}
