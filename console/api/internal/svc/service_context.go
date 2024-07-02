package svc

import (
	"4hfire/api/internal/config"
	"4hfire/api/internal/middleware"
	"4hfire/common/lib/jwt"
	"4hfire/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	Auth              rest.Middleware
	RedisClient       *redis.Redis
	MysqlConn         sqlx.SqlConn
	JwtInstance       jwt.JWT
	AdminsModel       model.AdminsModel
	AdminAvatarsModel model.AdminAvatarsModel
	CertsModel        model.CertsModel
	HostsModel        model.HostsModel
	HostCertsModel    model.HostCertsModel
	TagsModel         model.TagsModel
	RulesModel        model.RulesModel
	RuleGroupsModel   model.RuleGroupsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic("init redis client failed:" + err.Error())
	}
	JwtInstance := jwt.NewJWT(c.AuthCfg.AccessSecret, c.AuthCfg.AccessExpire, jwt.SetRedis(redisClient), jwt.SetBlackListOpt(true))
	srvCtx := &ServiceContext{
		Config:            c,
		RedisClient:       redisClient,
		MysqlConn:         conn,
		JwtInstance:       JwtInstance,
		Auth:              middleware.NewAuthMiddleware(JwtInstance).Handle,
		AdminsModel:       model.NewAdminsModel(conn, c.Mysql.DbCache),
		AdminAvatarsModel: model.NewAdminAvatarsModel(conn, c.Mysql.DbCache),
		CertsModel:        model.NewCertsModel(conn, c.Mysql.DbCache),
		HostsModel:        model.NewHostsModel(conn, c.Mysql.DbCache),
		HostCertsModel:    model.NewHostCertsModel(conn, c.Mysql.DbCache),
		TagsModel:         model.NewTagsModel(conn, c.Mysql.DbCache),
		RulesModel:        model.NewRulesModel(conn, c.Mysql.DbCache),
		RuleGroupsModel:   model.NewRuleGroupsModel(conn, c.Mysql.DbCache),
	}
	return srvCtx
}
