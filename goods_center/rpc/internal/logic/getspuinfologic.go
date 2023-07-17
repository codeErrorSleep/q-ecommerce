package logic

import (
	"context"
	"fmt"

	"goods_center/rpc/internal/svc"
	"goods_center/rpc/spu"

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
	// todo: add your logic here and delete this line
	fmt.Println("inside", in)

	if in.AppId == "1" {
		fmt.Println("dfsadfasd")
	}

	spuinfo := spu.SpuInfo{AppId: "dfsadaf", SpuId: "dfffffjj"}

	ret := spu.GetSpuInfoResponse{
		Products: []*spu.SpuInfo{&spuinfo},
	}

	return &ret, nil
}
