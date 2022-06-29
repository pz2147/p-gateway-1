package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/pz2147/p-gateway-1/internal/logic/user"
	"github.com/pz2147/p-gateway-1/internal/svc"
	"github.com/pz2147/p-gateway-1/internal/types"
)

func UserLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.GwHttpX.GwToApi(w, r, svc.ApiSvcType_User)
	}
}

func UserLogoutHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserLogoutLogic(r.Context(), ctx)
		resp, err := l.UserLogout(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.GwHttpX.GwToApi(w, r, svc.ApiSvcType_User)
	}
}



