package model

import (
	"4hfire/common/utils"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ AdminsModel = (*customAdminsModel)(nil)

const DEFAULT_ADMIN_USERNAME = "admin"
const DEFAULT_ADMIN_PASSWORD = "123456"

// admin账户的禁用启用状态
const (
	AdminIsNotDisabled = iota + 1 // 正常
	AdminIsDisabled               // 禁用
)

// OTP是否开启
const (
	AdminEnableOTP = iota + 1
	AdminDisableOTP
)

type (
	// AdminsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminsModel.
	AdminsModel interface {
		adminsModel
		CreateDefaultAdmin(ctx context.Context) error
		List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Admins, int64, error)
	}

	customAdminsModel struct {
		*defaultAdminsModel
	}
)

// NewAdminsModel returns a model for the database table.
func NewAdminsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AdminsModel {
	return &customAdminsModel{
		defaultAdminsModel: newAdminsModel(conn, c, opts...),
	}
}

func (m *defaultAdminsModel) CreateDefaultAdmin(ctx context.Context) error {
	t := time.Now().Unix()
	_, err := m.Insert(ctx, &Admins{
		Account:   DEFAULT_ADMIN_USERNAME,
		Password:  utils.MD5(utils.MD5(DEFAULT_ADMIN_PASSWORD)),
		CreatedAt: t,
		UpdatedAt: t,
		IsDisable: AdminIsNotDisabled,
		Email:     "",
		Uid:       uuid.NewV4().String(),
		Name:      "admin",
		Otp:       AdminEnableOTP,
	})
	return err
}
func (m *defaultAdminsModel) List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Admins, int64, error) {

	copyBuilder := sqlBuilder.Column("COUNT(*) as count").From(m.table)
	query, values, err := copyBuilder.ToSql()
	if err != nil {
		return nil, 0, err
	}

	// select count
	var count int64
	err = m.QueryRowNoCacheCtx(ctx, &count, query, values...)
	if err != nil {
		return nil, 0, err
	}

	var argFields []string
	for _, v := range adminsFieldNames {
		argFields = append(argFields, fmt.Sprintf("%s.%s", m.table, v))
	}

	// query data
	query, values, _ = sqlBuilder.Columns(argFields...).From(m.table).
		OrderBy("createdAt DESC").Offset(uint64((page - 1) * size)).Limit(uint64(size)).ToSql()

	var us []*Admins
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, 0, err
	}
	return us, count, nil
}
