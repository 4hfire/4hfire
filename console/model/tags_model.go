package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TagsModel = (*customTagsModel)(nil)

type (
	// TagsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagsModel.
	TagsModel interface {
		tagsModel
		List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Tags, int64, error)
		All(ctx context.Context) ([]*Tags, error)
	}

	customTagsModel struct {
		*defaultTagsModel
	}
)

// NewTagsModel returns a model for the database table.
func NewTagsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TagsModel {
	return &customTagsModel{
		defaultTagsModel: newTagsModel(conn, c, opts...),
	}
}

func (m *defaultTagsModel) All(ctx context.Context) ([]*Tags, error) {
	query, values, err := squirrel.Select().From(m.table).ToSql()
	if err != nil {
		return nil, err
	}
	var us []*Tags
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, err
	}
	return us, nil
}
func (m *defaultTagsModel) List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Tags, int64, error) {

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

	var us []*Tags
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, 0, err
	}
	return us, count, nil
}
