package handler

import (
	"net/http"

	"EthTest/demo/internal/logic"
	"EthTest/demo/internal/svc"
	"EthTest/demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ERC20TransferHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ERC20TransferRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewERC20TransferLogic(r.Context(), svcCtx)
		resp, err := l.ERC20Transfer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
