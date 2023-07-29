package spu

import (
	"context"

	"goods_center/api/internal/svc"
	"goods_center/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetSpuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuListLogic {
	return &GetSpuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSpuListLogic) GetSpuList(req *types.GetSpuInfoReq) (resp *types.GetSpuInfoResp, err error) {

	resp = new(types.GetSpuInfoResp)
	resp.Data.SpuInfo.AppId = "1211"

	return resp, status.Error(110, "dsfadsfas")
}
