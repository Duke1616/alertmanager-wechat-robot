package endpoint

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/types/ftime"
	"hash/fnv"
)

const (
	AppName = "endpoint"
)

func (req *RegistryRequest) Endpoints() []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			Id:       GenHashID("alertmanager-wechat-robot", req.Entries[i].Path),
			CreateAt: ftime.Now().Timestamp(),
			UpdateAt: ftime.Now().Timestamp(),
			Version:  req.Version,
			Entry:    req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

func GenHashID(service, grpcPath string) string {
	hashedStr := fmt.Sprintf("%s-%s", service, grpcPath)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

func TransferRoutesToEntry(routes []restful.Route) (entries []*Entry) {
	for _, r := range routes {
		entries = append(entries, NewEntryFromRestRoute(r))
	}
	return
}

func NewDefaultEntry() *Entry {
	return &Entry{}
}

func NewEntryFromRestRoute(route restful.Route) *Entry {
	entry := NewDefaultEntry()
	entry.FunctionName = route.Operation
	entry.Method = route.Method
	entry.LoadMeta(route.Metadata)
	entry.Path = route.Path

	entry.Path = entry.UniquePath()
	return entry
}

func (e *Entry) UniquePath() string {
	return fmt.Sprintf("%s.%s", e.Method, e.Path)
}

func (e *Entry) LoadMeta(meta map[string]interface{}) {
	if meta != nil {
		if v, ok := meta[label.Resource]; ok {
			e.Resource, _ = v.(string)
		}
		if v, ok := meta[label.Auth]; ok {
			e.AuthEnable, _ = v.(bool)
		}
	}
}

// NewRegistryRequest 注册请求
func NewRegistryRequest(version string, entries []*Entry) *RegistryRequest {
	return &RegistryRequest{
		Version: version,
		Entries: entries,
	}
}

func NewEntryFromRestRequest(req *restful.Request) *Entry {
	entry := NewDefaultEntry()

	// 请求拦截
	route := req.SelectedRoute()
	if route == nil {
		return nil
	}

	entry.FunctionName = route.Operation()
	entry.Method = route.Method()
	entry.LoadMeta(route.Metadata())
	entry.Path = route.Path()

	entry.Path = entry.UniquePath()
	return entry
}
