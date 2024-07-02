package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CertsModel = (*customCertsModel)(nil)

type (
	// CertsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCertsModel.
	CertsModel interface {
		certsModel
		List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Certs, int64, error)
		All(ctx context.Context) ([]*Certs, error)
		Count(ctx context.Context, sqlBuilder squirrel.SelectBuilder) (uint64, error)

		InsertSession(ctx context.Context, data *Certs, s sqlx.Session) (int64, error)
	}

	customCertsModel struct {
		*defaultCertsModel
	}
)

// NewCertsModel returns a model for the database table.
func NewCertsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CertsModel {
	return &customCertsModel{
		defaultCertsModel: newCertsModel(conn, c, opts...),
	}
}

func (m *defaultCertsModel) Count(ctx context.Context, sqlBuilder squirrel.SelectBuilder) (uint64, error) {
	copyBuilder := sqlBuilder.Column("COUNT(*) as count").From(m.table)
	query, values, err := copyBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	// select count
	var count uint64
	err = m.QueryRowNoCacheCtx(ctx, &count, query, values...)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (m *defaultCertsModel) All(ctx context.Context) ([]*Certs, error) {
	query, values, err := squirrel.Select().From(m.table).ToSql()
	if err != nil {
		return nil, err
	}
	var us []*Certs
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, err
	}
	return us, nil
}
func (m *defaultCertsModel) List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Certs, int64, error) {

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

	var us []*Certs
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, 0, err
	}
	return us, count, nil
}

func (m *defaultCertsModel) InsertSession(ctx context.Context, data *Certs, s sqlx.Session) (int64, error) {

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, certsRowsExpectAutoSet)
	res, err := s.ExecCtx(ctx, query, data.Name, data.Account, data.Password, data.Secret, data.Desc, data.Type)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
