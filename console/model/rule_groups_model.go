package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RuleGroupsModel = (*customRuleGroupsModel)(nil)

type (
	// RuleGroupsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRuleGroupsModel.
	RuleGroupsModel interface {
		ruleGroupsModel
		List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*RuleGroups, int64, error)
		All(ctx context.Context) ([]*RuleGroups, error)
		Count(ctx context.Context, sqlBuilder squirrel.SelectBuilder) (uint64, error)
	}

	customRuleGroupsModel struct {
		*defaultRuleGroupsModel
	}
)

// NewRuleGroupsModel returns a model for the database table.
func NewRuleGroupsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RuleGroupsModel {
	return &customRuleGroupsModel{
		defaultRuleGroupsModel: newRuleGroupsModel(conn, c, opts...),
	}
}

func (m *defaultRuleGroupsModel) Count(ctx context.Context, sqlBuilder squirrel.SelectBuilder) (uint64, error) {
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
func (m *defaultRuleGroupsModel) All(ctx context.Context) ([]*RuleGroups, error) {
	query, values, err := squirrel.Select().From(m.table).ToSql()
	if err != nil {
		return nil, err
	}
	var us []*RuleGroups
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, err
	}
	return us, nil
}
func (m *defaultRuleGroupsModel) List(ctx context.Context, sqlBuilder squirrel.SelectBuilder, page, size int64) ([]*RuleGroups, int64, error) {

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

	var us []*RuleGroups
	err = m.QueryRowsNoCacheCtx(ctx, &us, query, values...)
	if err != nil {
		return nil, 0, err
	}
	return us, count, nil
}
