package rule

import (
	"4hfire/common/errors"
	"4hfire/common/utils"
	"4hfire/model"
	"context"
	"strconv"
	"strings"
	"unicode/utf8"

	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.RuleAddReq) error {
	if _, err := l.svcCtx.RulesModel.FindOneByName(l.ctx, req.Name); err != model.ErrNotFound {
		logx.Error(err)
		err = errors.RuleAddRepatedError()
		return err
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
	_, err := l.svcCtx.RulesModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return errors.DbError()
	}
	return nil
}

// 1;1-3;2
func checkPort(ports string) bool {
	for _, vv := range strings.Split(ports, ";") {
		tmp := strings.Split(vv, "-")
		if len(tmp) > 2 {
			return false
		}
		ps := [2]int{}
		for k, v := range tmp {
			ps[k], _ = strconv.Atoi(v)
			if ps[k] > 65535 && ps[k] < 1 {
				return false
			}
		}
		if ps[1] > 0 && ps[0] > ps[1] {
			return false
		}
	}
	return false
}
