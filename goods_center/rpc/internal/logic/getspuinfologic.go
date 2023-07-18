package logic

import (
	"context"

	"goods_center/rpc/internal/svc"
	"goods_center/rpc/spu"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpuInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuInfoLogic {
	return &GetSpuInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSpuInfoLogic) GetSpuInfo(in *spu.GetSpuInfoRequest) (*spu.GetSpuInfoResponse, error) {
	if in == nil || in.AppId == "" {
		logc.Info(l.ctx, "input wrong", in)
	}
	// time.Sleep(time.Second * 2)
	spuModel, err := l.svcCtx.SpuModel.FindOneByAppIdSpuId(l.ctx, in.AppId, in.SpuId)
	if err != nil {
		logc.Info(l.ctx, "find spu error")
	}

	ret := spu.GetSpuInfoResponse{
		SpuInfo: &spu.SpuInfo{},
	}

	if spuModel == nil {
		return &ret, nil
	}

	ret = spu.GetSpuInfoResponse{
		SpuInfo: &spu.SpuInfo{
			AppId: spuModel.AppId,
			SpuId: spuModel.SpuId,
		},
	}

	return &ret, nil
}
