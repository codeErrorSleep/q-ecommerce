package svc

import (
	"goods_center/rpc/internal/config"
	"goods_center/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	SpuModel model.TSpuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:   c,
		SpuModel: model.NewTSpuModel(conn),
	}
}
