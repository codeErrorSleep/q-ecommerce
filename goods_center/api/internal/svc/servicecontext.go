package svc

import (
	"goods_center/api/internal/config"
	"goods_center/rpc/spuclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	RPC    spuclient.Spu
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RPC:    spuclient.NewSpu(zrpc.MustNewClient(c.RPC)),
	}
}
