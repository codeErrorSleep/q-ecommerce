package spu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goods_center/api/internal/logic/spu"
	"goods_center/api/internal/svc"
	"goods_center/api/internal/types"
)

func GetSpuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSpuInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := spu.NewGetSpuListLogic(r.Context(), svcCtx)
		resp, err := l.GetSpuList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
