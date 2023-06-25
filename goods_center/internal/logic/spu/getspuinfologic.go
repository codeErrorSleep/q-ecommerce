package spu

import (
	"context"

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
	// todo: add your logic here and delete this line
	resp = new(types.GetSpuInfoResp)
	resp.AppId = "dfassdfa"
	return
}