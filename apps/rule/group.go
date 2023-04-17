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

func NewGroupSet() *GroupSet {
	return &GroupSet{
		Items: []*Group{},
	}
}

func (s *GroupSet) Add(item *Group) {
	s.Total++
	s.Items = append(s.Items, item)
}

func NewDefaultGroup() *Group {
	return &Group{}
}
