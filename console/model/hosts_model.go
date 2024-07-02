package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HostsModel = (*customHostsModel)(nil)

type (
	// HostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHostsModel.
	HostsModel interface {
		hostsModel
		List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Hosts, int64, error)
		InsertSession(ctx context.Context, data *Hosts, s sqlx.Session) (int64, error)
		DeleteSession(ctx context.Context, id uint64, s sqlx.Session) error
	}

	customHostsModel struct {
		*defaultHostsModel
	}
)

// NewHostsModel returns a model for the database table.
func NewHostsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HostsModel {
	return &customHostsModel{
		defaultHostsModel: newHostsModel(conn, c, opts...),
	}
}

func (m *defaultHostsModel) InsertSession(ctx context.Context, data *Hosts, s sqlx.Session) (int64, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, hostsRowsExpectAutoSet)
	res, err := s.ExecCtx(ctx, query, data.Id, data.Uuid, data.Secret, data.Address, data.Tags, data.Name, data.Desc, data.State, data.Cpu, data.Mem)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func (m *defaultHostsModel) DeleteSession(ctx context.Context, id uint64, s sqlx.Session) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	hostsIdKey := fmt.Sprintf("%s%v", cacheHostsIdPrefix, id)
	hostsUuidKey := fmt.Sprintf("%s%v", cacheHostsUuidPrefix, data.Uuid)
	err = m.DelCacheCtx(ctx, hostsIdKey, hostsUuidKey)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err = s.ExecCtx(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultHostsModel) List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Hosts, int64, error) {

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

	var us []*Hosts
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, 0, err
	}
	return us, count, nil
}
