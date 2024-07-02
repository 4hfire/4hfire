package host

import (
	"4hfire/common/errors"
	"4hfire/common/installer"
	"4hfire/common/utils"
	"4hfire/model"
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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

func (l *AddLogic) Add(req *types.HostAddReq) error {
	if _, err := l.svcCtx.CertsModel.FindOneByName(l.ctx, req.Name); err != model.ErrNotFound {
		logx.Error(err)
		err = errors.RuleAddRepatedError()
		return err
	}
	if utf8.RuneCountInString(req.Name) > 32 {
		return errors.ParamsError("主机名称长度不合规，长度需要小于64位")
	}
	if utf8.RuneCountInString(req.Desc) > 255 {
		return errors.ParamsError("主机备注长度不合规，长度需要小于255位")
	}

	data := &model.Hosts{
		Name:    req.Name,
		Desc:    req.Desc,
		Tags:    utils.ToJSONString(req.Tags),
		Uuid:    strings.Trim(uuid.NewV4().String(), "-"),
		Secret:  strings.Trim(uuid.NewV4().String(), "-"),
		Address: req.Address,
	}

	if req.AutoInstall {
		if req.CertId == 0 && req.AutoCert { // 自动同步凭证
			if utf8.RuneCountInString(req.Account) > 64 {
				return errors.ParamsError("凭证账号长度不合规，长度需要小于255位")
			}
			if utf8.RuneCountInString(req.Password) > 64 {
				return errors.ParamsError("凭证密码长度不合规，长度需要小于255位")
			}
			if utf8.RuneCountInString(req.Secret) > 1024 {
				return errors.ParamsError("凭证私钥长度不合规，长度需要小于255位")
			}
			if len(req.Account) == 0 && ((req.Type == 1 && len(req.Password) == 0) ||
				(req.Type == 2 && len(req.Secret) == 0)) {
				return errors.ParamsError()
			}
		} else if req.CertId != 0 {
			cert, err := l.svcCtx.CertsModel.FindOne(l.ctx, req.CertId)
			if err != nil {
				logx.Error(err)
				err = errors.DbError()
				return err
			}
			req.Account = cert.Account
			req.Password = cert.Password
			req.Secret = cert.Secret
		}
		cert := &model.Certs{
			Name:     req.Name + "主机(自动同步)",
			Desc:     "主机：" + req.Name + "，自动同步凭证",
			Account:  req.Account,
			Password: req.Password,
			Secret:   req.Secret,
			Type:     req.Type,
		}
		err := l.svcCtx.MysqlConn.TransactCtx(l.ctx, func(ctx context.Context, s sqlx.Session) error {
			if req.CertId == 0 && req.AutoCert {
				res, err := l.svcCtx.CertsModel.InsertSession(l.ctx, cert, s)
				if err != nil {
					logx.Error(err)
					return errors.DbError()
				}
				req.CertId = uint64(res)
			}
			hc := &model.HostCerts{CertId: req.CertId, Account: req.Account, Password: req.Password, Secret: req.Secret, Type: req.Type}
			res, err := l.svcCtx.HostsModel.InsertSession(l.ctx, data, s)
			if err != nil {
				logx.Error(err)
				return errors.DbError()
			}
			hc.Id = uint64(res)
			_, err = l.svcCtx.HostCertsModel.InsertSession(l.ctx, hc, s)
			return err
		})
		if err != nil {
			logx.Error(err)
			return errors.SystemError()
		}

		// todo：实现SSH远程进入服务器执行安装命令

		install := installer.NewInstaller()
		err = install.Login(&installer.Credentials{
			Host:       req.IP,
			Port:       req.Port,
			Username:   req.Account,
			Password:   req.Password,
			PrivateKey: req.Secret,
			Passphrase: req.Password,
		})
		if err != nil {
			err = errors.DiyError("资产连接失败")
			return err
		}
		err = install.Install(map[string]string{
			"endpoint": l.svcCtx.Config.Host,
			"nodeId":   data.Uuid,
			"secret":   data.Secret,
		})
		if err != nil {
			err = errors.DiyError(err.Error())
			return err
		}
	} else {
		_, err := l.svcCtx.HostsModel.Insert(l.ctx, data)
		if err != nil {
			logx.Error(err)
			return errors.DbError()
		}
	}
	return nil
}
