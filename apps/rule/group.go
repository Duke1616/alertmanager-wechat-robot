package rule

import (
	"github.com/infraboard/mcube/http/request"
	"net/http"
)

func (req *QueryGroupRequest) Validate() error {
	return validate.Struct(req)
}

func NewQueryGroupRequestFromHTTP(r *http.Request) *QueryGroupRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryGroupRequest()
	req.Name = qs.Get("name")

	req.Page = page
	return req
}

func NewDefaultQueryGroupRequest() *QueryGroupRequest {
	return &QueryGroupRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
