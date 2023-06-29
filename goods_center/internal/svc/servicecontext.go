package svc

import (
	"goods_center/internal/config"
	"goods_center/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	SpuModel model.TSpuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config:   c,
		SpuModel: model.NewTSpuModel(sqlConn),
	}
}
