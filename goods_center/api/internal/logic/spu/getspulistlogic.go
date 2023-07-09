package spu

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return
}
