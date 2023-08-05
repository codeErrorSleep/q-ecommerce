package sql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TOrdersModel = (*customTOrdersModel)(nil)

type (
	// TOrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTOrdersModel.
	TOrdersModel interface {
		tOrdersModel
	}

	customTOrdersModel struct {
		*defaultTOrdersModel
	}
)

// NewTOrdersModel returns a model for the database table.
func NewTOrdersModel(conn sqlx.SqlConn) TOrdersModel {
	return &customTOrdersModel{
		defaultTOrdersModel: newTOrdersModel(conn),
	}
}
