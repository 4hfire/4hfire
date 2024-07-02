package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HostCertsModel = (*customHostCertsModel)(nil)

type (
	// HostCertsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHostCertsModel.
	HostCertsModel interface {
		hostCertsModel
		InsertSession(ctx context.Context, data *HostCerts, s sqlx.Session) (int64, error)
		DeleteSession(ctx context.Context, id uint64, s sqlx.Session) error
	}

	customHostCertsModel struct {
		*defaultHostCertsModel
	}
)

// NewHostCertsModel returns a model for the database table.
func NewHostCertsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HostCertsModel {
	return &customHostCertsModel{
		defaultHostCertsModel: newHostCertsModel(conn, c, opts...),
	}
}

func (m *defaultHostCertsModel) InsertSession(ctx context.Context, data *HostCerts, s sqlx.Session) (int64, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, hostCertsRowsExpectAutoSet)
	res, err := s.ExecCtx(ctx, query, data.Id, data.Account, data.Password, data.Secret, data.CertId, data.Type)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (m *defaultHostCertsModel) DeleteSession(ctx context.Context, id uint64, s sqlx.Session) error {

	hostsIdKey := fmt.Sprintf("%s%v", cacheHostsIdPrefix, id)
	err := m.DelCacheCtx(ctx, hostsIdKey)
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
