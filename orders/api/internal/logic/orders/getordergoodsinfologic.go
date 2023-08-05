package orders

import (
	"context"

	"orders/api/internal/svc"
	"orders/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderGoodsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderGoodsInfoLogic {
	return &GetOrderGoodsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderGoodsInfoLogic) GetOrderGoodsInfo(req *types.GetOrderGoodsInfoReq) (resp *types.GetOrderGoodsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
