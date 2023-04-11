package notifier

import (
	"net/http"
)

const (
	AppName = "notifier"
)

func NewQueryTargetRequestFromHTTP(r *http.Request, name string) string {
	qs := r.URL.Query()
	return qs.Get(name)
}
