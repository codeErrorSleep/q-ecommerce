package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

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

func (m *defaultTSpuModel) FindOneByAppIdSpuId(ctx context.Context, appId string, spuId string) (*TSpu, error) {
	var resp TSpu
	query := fmt.Sprintf("select %s from %s where `app_id` = ? and `spu_id` = ? limit 1", tSpuRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, appId, spuId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
