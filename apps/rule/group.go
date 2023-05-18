package rule

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"github.com/infraboard/mcube/http/request"
	"net/http"
)

func (req *QueryGroupRequest) Validate() error {
	return validate.Struct(req)
}

func (req *Rules) GroupSet(setting *setting.Setting) []interface{} {
	gps := make([]interface{}, 0, len(req.Data.Groups))
	for i := range req.Data.Groups {
		gp := &Group{
			Id:        req.Data.Groups[i].Id,
			Name:      req.Data.Groups[i].Name,
			Domain:    setting.Spec.Domain,
			Namespace: setting.Spec.Namespace,
			Interval:  req.Data.Groups[i].Interval,
		}

		gps = append(gps, gp)
	}
	return gps
}

func NewQueryGroupRequestFromHTTP(r *http.Request) *QueryGroupRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryGroupRequest()
	req.Domain = qs.Get("domain")
	req.Namespace = qs.Get("namespace")

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

func NewDeleteGroupRequest(domain, namespace string) *DeleteGroupRequest {
	return &DeleteGroupRequest{
		Domain:    domain,
		Namespace: namespace,
	}
}

func (s *GroupSet) GroupsIds() (gps []string) {
	for i := range s.Items {
		gps = append(gps, s.Items[i].Id)
	}

	return
}
