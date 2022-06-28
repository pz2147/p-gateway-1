package handler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"github.com/pz2147/p-gateway-1/internal/logic/user"
	"github.com/pz2147/p-gateway-1/internal/svc"
	"github.com/pz2147/p-gateway-1/internal/types"

	"github.com/imroc/req"
)

func UserLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//var reqMap map[string]interface{}
		//if err := httpx.ParseJsonBody(r, &reqMap); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}

		url := ctx.Config.ForwardConf.Api1.Host + r.RequestURI

		gReq := req.New()
		if r.Method == http.MethodPost {
			gResp, err := gReq.Post(url, r.Header, r.Body)
			if err != nil {
				httpx.Error(w, err)
			} else {
				var respMap map[string]interface{}
				respErr := gResp.ToJSON(&respMap)
				if respErr != nil {
					httpx.Error(w, err)
				} else {
					httpx.OkJson(w, respMap)
				}
			}
		}

		//l := user.NewUserLoginLogic(r.Context(), ctx)
		//resp, err := l.UserLogin(req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
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
		var req types.EmptyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
