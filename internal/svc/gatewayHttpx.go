package svc

import (
	"errors"
	"github.com/imroc/req"
	"github.com/pz2147/p-gateway-1/internal/common/util"
	"github.com/pz2147/p-gateway-1/internal/config"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type ApiSvcType int32

const (
	ApiSvcType_None ApiSvcType = 0 //无服务
	ApiSvcType_User ApiSvcType = 1 //用户服务
)

type GatewayHttpX struct {
	Config config.Config
}

func NewGatewayHttpX(c config.Config) *GatewayHttpX {
	return &GatewayHttpX{
		Config: c,
	}
}

// GwToApi gateway 转发到api
func (x *GatewayHttpX) GwToApi(w http.ResponseWriter, r *http.Request, asType ApiSvcType) {

	// 1. 配置TraceId
	x.setUpTraceId(r)

	// 2. 转发的链接
	url := x.getHost(asType) + r.RequestURI

	gReq := req.New()
	if r.Method == http.MethodPost {
		gResp, err := gReq.Post(url, r.Header, r.Body)
		if err != nil {
			httpx.Error(w, err)
		} else {
			var respMap map[string]interface{}
			respErr := gResp.ToJSON(&respMap)
			if respErr != nil {
				httpx.Error(w, respErr)
			} else {
				httpx.OkJson(w, respMap)
			}
		}
	} else if r.Method == http.MethodGet {
		gResp, err := gReq.Get(url, r.Header, r.Body)
		if err != nil {
			httpx.Error(w, err)
		} else {
			if gResp.Response().StatusCode != 200 {
				httpx.Error(w, errors.New(gResp.String()))
			} else {
				var respMap map[string]interface{}
				respErr := gResp.ToJSON(&respMap)
				if respErr != nil {
					httpx.Error(w, respErr)
				} else {
					httpx.OkJson(w, respMap)
				}
			}
		}
	}
}

// setUpTraceId 配置Traceid
func (x *GatewayHttpX) setUpTraceId(r *http.Request) {
	var su util.SpanUtil
	str := su.GetTraceParentStr(r.Context())
	if len(str) > 0 {
		r.Header.Set(util.Traceparent, str)
	}
}

// getHost 获取host
func (x *GatewayHttpX) getHost(asType ApiSvcType) string {

	host := ""

	conf := x.Config.ForwardConf
	switch asType {
	case ApiSvcType_User:
		host = conf.Api1.Host
		break
	default:
		break
	}
	return host
}
