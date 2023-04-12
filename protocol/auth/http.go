package auth

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/endpoint"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/token"
	"github.com/Duke1616/alertmanager-wechat-robot/conf"
	app "github.com/Duke1616/alertmanager-wechat-robot/register"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewHttpAuthor() *httpAuthor {
	return &httpAuthor{
		log:  zap.L().Named("author.http"),
		tk:   app.GetInternalApp(token.AppName).(token.RPCServer),
		conf: conf.C(),
	}
}

type httpAuthor struct {
	log  logger.Logger
	tk   token.RPCServer
	conf *conf.Config
}

func (a *httpAuthor) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截
	entry := endpoint.NewEntryFromRestRequest(req)

	if entry != nil && entry.AuthEnable {
		// 访问令牌校验
		_, err := a.CheckAccessToken(req)
		if err != nil {
			response.Failed(resp, err)
			return
		}
	}

	next.ProcessFilter(req, resp)
}

func (a *httpAuthor) CheckAccessToken(req *restful.Request) (*token.Token, error) {
	ak := token.GetAccessTokenFromHTTP(req.Request)

	if ak == "" {
		return nil, exception.NewUnauthorized("Auth Header Required, Format: %s: Bearer ${access_token}", token.ACCESS_TOKEN_HEADER_KEY)
	}

	// 调用GRPC 校验用户Token合法性
	tk, err := a.tk.ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
	if err != nil {
		return nil, err
	}

	req.SetAttribute(token.TOKEN_ATTRIBUTE_NAME, tk)
	return tk, nil
}
