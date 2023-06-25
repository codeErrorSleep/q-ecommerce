package spu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goods_center/internal/logic/spu"
	"goods_center/internal/svc"
	"goods_center/internal/types"
)

func GetSpuInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSpuInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := spu.NewGetSpuInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetSpuInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
