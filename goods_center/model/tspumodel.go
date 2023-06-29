package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TSpuModel = (*customTSpuModel)(nil)

type (
	// TSpuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSpuModel.
	TSpuModel interface {
		tSpuModel
	}

	customTSpuModel struct {
		*defaultTSpuModel
	}
)

// NewTSpuModel returns a model for the database table.
func NewTSpuModel(conn sqlx.SqlConn) TSpuModel {
	return &customTSpuModel{
		defaultTSpuModel: newTSpuModel(conn),
	}
}
