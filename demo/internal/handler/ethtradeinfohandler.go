package handler

import (
	"net/http"

	"EthTest/demo/internal/logic"
	"EthTest/demo/internal/svc"
	"EthTest/demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EthTradeInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TradeInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewEthTradeInfoLogic(r.Context(), svcCtx)
		resp, err := l.EthTradeInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
