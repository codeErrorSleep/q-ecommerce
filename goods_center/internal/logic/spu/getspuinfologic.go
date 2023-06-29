package spu

import (
	"context"
	"fmt"

	"goods_center/internal/svc"
	"goods_center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	spuData, err := l.svcCtx.SpuModel.FindOne(l.ctx, 1)
	if err != nil {
		fmt.Println(err)
	}

	resp = new(types.GetSpuInfoResp)
	resp.AppId = spuData.AppId
	resp.GoodsName = spuData.GoodsName

	return
}
