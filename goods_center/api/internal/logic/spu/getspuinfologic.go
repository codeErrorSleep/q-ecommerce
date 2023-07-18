package spu

import (
	"context"

	"goods_center/api/internal/svc"
	"goods_center/api/internal/types"
	"goods_center/rpc/spu"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetSpuInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuInfoLogic {
	return &GetSpuInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSpuInfoLogic) GetSpuInfo(req *types.GetSpuInfoReq) (resp *types.GetSpuInfoResp, err error) {
	spuInfo, err := l.svcCtx.RPC.GetSpuInfo(l.ctx, &spu.GetSpuInfoRequest{AppId: req.AppId, SpuId: req.SpuId})
	if err != nil {
		return nil, status.Error(500, "查表异常")
	}
	resp = new(types.GetSpuInfoResp)

	spuData := types.Spu{
		AppId:     spuInfo.SpuInfo.AppId,
		SpuId:     spuInfo.SpuInfo.SpuId,
		GoodsName: spuInfo.SpuInfo.GoodsName,
	}
	resp.Data.SpuInfo = spuData

	return resp, nil
}
