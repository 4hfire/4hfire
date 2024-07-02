// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tagsFieldNames          = builder.RawFieldNames(&Tags{})
	tagsRows                = strings.Join(tagsFieldNames, ",")
	tagsRowsExpectAutoSet   = strings.Join(stringx.Remove(tagsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tagsRowsWithPlaceHolder = strings.Join(stringx.Remove(tagsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTagsIdPrefix   = "cache:tags:id:"
	cacheTagsNamePrefix = "cache:tags:name:"
)

type (
	tagsModel interface {
		Insert(ctx context.Context, data *Tags) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Tags, error)
		FindOneByName(ctx context.Context, name string) (*Tags, error)
		Update(ctx context.Context, data *Tags) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultTagsModel struct {
		sqlc.CachedConn
		table string
	}

	Tags struct {
		Id   uint64 `db:"id"`   // 标签id
		Name string `db:"name"` // 名称
		Desc string `db:"desc"` // 备注
	}
)

func newTagsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTagsModel {
	return &defaultTagsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`tags`",
	}
}

func (m *defaultTagsModel) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	tagsIdKey := fmt.Sprintf("%s%v", cacheTagsIdPrefix, id)
	tagsNameKey := fmt.Sprintf("%s%v", cacheTagsNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tagsIdKey, tagsNameKey)
	return err
}

func (m *defaultTagsModel) FindOne(ctx context.Context, id uint64) (*Tags, error) {
	tagsIdKey := fmt.Sprintf("%s%v", cacheTagsIdPrefix, id)
	var resp Tags
	err := m.QueryRowCtx(ctx, &resp, tagsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tagsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTagsModel) FindOneByName(ctx context.Context, name string) (*Tags, error) {
	tagsNameKey := fmt.Sprintf("%s%v", cacheTagsNamePrefix, name)
	var resp Tags
	err := m.QueryRowIndexCtx(ctx, &resp, tagsNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", tagsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTagsModel) Insert(ctx context.Context, data *Tags) (sql.Result, error) {
	tagsIdKey := fmt.Sprintf("%s%v", cacheTagsIdPrefix, data.Id)
	tagsNameKey := fmt.Sprintf("%s%v", cacheTagsNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, tagsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Desc)
	}, tagsIdKey, tagsNameKey)
	return ret, err
}

func (m *defaultTagsModel) Update(ctx context.Context, newData *Tags) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	tagsIdKey := fmt.Sprintf("%s%v", cacheTagsIdPrefix, data.Id)
	tagsNameKey := fmt.Sprintf("%s%v", cacheTagsNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tagsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Desc, newData.Id)
	}, tagsIdKey, tagsNameKey)
	return err
}

func (m *defaultTagsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTagsIdPrefix, primary)
}

func (m *defaultTagsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tagsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTagsModel) tableName() string {
	return m.table
}