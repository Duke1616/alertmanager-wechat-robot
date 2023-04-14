package utils

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
)

func Success(w *restful.Response, data interface{}, opts ...response.Option) {
	var (
		httpCode int
	)

	switch t := data.(type) {
	case exception.APIException:
		data = t.GetData()
		httpCode = t.GetHttpCode()
	}

	if httpCode == 0 {
		httpCode = http.StatusOK
	}

	resp := response.Data{
		Code:    &httpCode,
		Message: "成功",
		Data:    data,
	}

	for _, opt := range opts {
		opt.Apply(&resp)
	}

	err := w.WriteHeaderAndEntity(httpCode, resp)
	if err != nil {
		zap.L().Errorf("send failed response error, %s", err)
	}
}
