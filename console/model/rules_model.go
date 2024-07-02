package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RulesModel = (*customRulesModel)(nil)

// rule的禁用启用状态
const (
	RuleDisabled = iota // 禁用
	RuleEnable          // 正常
)

// rule的执行动作
const (
	RuleOptionAccept = iota // 允许
	RuleOptionDrop          // 阻止
)

type (
	// RulesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRulesModel.
	RulesModel interface {
		rulesModel

		List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Rules, int64, error)
		All(ctx context.Context) ([]*Rules, error)
		Count(ctx context.Context, sqlBuilder squirrel.SelectBuilder) (uint64, error)
	}

	customRulesModel struct {
		*defaultRulesModel
	}
)

// NewRulesModel returns a model for the database table.
func NewRulesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RulesModel {
	return &customRulesModel{
		defaultRulesModel: newRulesModel(conn, c, opts...),
	}
}

func (m *defaultRulesModel) Count(ctx context.Context, sqlBuilder squirrel.SelectBuilder) (uint64, error) {
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
func (m *defaultRulesModel) All(ctx context.Context) ([]*Rules, error) {
	query, values, err := squirrel.Select().From(m.table).ToSql()
	if err != nil {
		return nil, err
	}
	var us []*Rules
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, err
	}
	return us, nil
}
func (m *defaultRulesModel) List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*Rules, int64, error) {

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

	var us []*Rules
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, 0, err
	}
	return us, count, nil
}
