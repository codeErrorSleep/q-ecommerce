// Code generated by goctl. DO NOT EDIT.

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tOrdersFieldNames          = builder.RawFieldNames(&TOrders{})
	tOrdersRows                = strings.Join(tOrdersFieldNames, ",")
	tOrdersRowsExpectAutoSet   = strings.Join(stringx.Remove(tOrdersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tOrdersRowsWithPlaceHolder = strings.Join(stringx.Remove(tOrdersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tOrdersModel interface {
		Insert(ctx context.Context, data *TOrders) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TOrders, error)
		Update(ctx context.Context, data *TOrders) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTOrdersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TOrders struct {
		Id        int64     `db:"id"`         // 自增id
		ShopId    string    `db:"shop_id"`    // 店铺id
		UserId    string    `db:"user_id"`    // 用户id
		OrderId   string    `db:"order_id"`   // 订单id
		GoodsId   string    `db:"goods_id"`   // 商品id
		Money     int64     `db:"money"`      // 订单金额
		Num       int64     `db:"num"`        // 购买数量
		Status    int64     `db:"status"`     // 订单状态 1:待支付,2:已支付,3:待退款,4:已退款
		IsDeleted int64     `db:"is_deleted"` // 是否删除：0-正常，1-删除
		CreatedAt time.Time `db:"created_at"` // 创建时间
		UpdatedAt time.Time `db:"updated_at"` // 更新时间
	}
)

func newTOrdersModel(conn sqlx.SqlConn) *defaultTOrdersModel {
	return &defaultTOrdersModel{
		conn:  conn,
		table: "`t_orders`",
	}
}

func (m *defaultTOrdersModel) withSession(session sqlx.Session) *defaultTOrdersModel {
	return &defaultTOrdersModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`t_orders`",
	}
}

func (m *defaultTOrdersModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTOrdersModel) FindOne(ctx context.Context, id int64) (*TOrders, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tOrdersRows, m.table)
	var resp TOrders
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTOrdersModel) Insert(ctx context.Context, data *TOrders) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, tOrdersRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ShopId, data.UserId, data.OrderId, data.GoodsId, data.Money, data.Num, data.Status, data.IsDeleted)
	return ret, err
}

func (m *defaultTOrdersModel) Update(ctx context.Context, data *TOrders) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tOrdersRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ShopId, data.UserId, data.OrderId, data.GoodsId, data.Money, data.Num, data.Status, data.IsDeleted, data.Id)
	return err
}

func (m *defaultTOrdersModel) tableName() string {
	return m.table
}