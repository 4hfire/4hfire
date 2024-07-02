package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminAvatarsModel = (*customAdminAvatarsModel)(nil)

type (
	// AdminAvatarsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminAvatarsModel.
	AdminAvatarsModel interface {
		adminAvatarsModel
	}

	customAdminAvatarsModel struct {
		*defaultAdminAvatarsModel
	}
)

// NewAdminAvatarsModel returns a model for the database table.
func NewAdminAvatarsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AdminAvatarsModel {
	return &customAdminAvatarsModel{
		defaultAdminAvatarsModel: newAdminAvatarsModel(conn, c, opts...),
	}
}
