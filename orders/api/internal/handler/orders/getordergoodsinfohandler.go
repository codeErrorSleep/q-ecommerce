package orders

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orders/api/internal/logic/orders"
	"orders/api/internal/svc"
	"orders/api/internal/types"
)

func GetOrderGoodsInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderGoodsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := orders.NewGetOrderGoodsInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderGoodsInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
