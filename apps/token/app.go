package token

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/logger/zap"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "token"
)

const (
	ACCESS_TOKEN_HEADER_KEY = "Authorization"
	ACCESS_TOKEN_COOKIE_KEY = "alertmanager.access_token"
	TOKEN_ATTRIBUTE_NAME    = "token"
)

var (
	validate = validator.New()
)

func (req *IssueTokenRequest) Validate() error {
	return validate.Struct(req)
}

func NewToken(req *IssueTokenRequest) (*Token, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	tk := &Token{
		AccessToken:     MakeBearer(24),
		IssueAt:         time.Now().UnixMilli(),
		AccessExpiredAt: 3600,
		Platform:        PLATFORM_WEB,
	}
	return tk, nil
}

func MakeBearer(length int) string {
	char := "123456"
	t := make([]string, length)
	rand.Seed(time.Now().UnixNano() + int64(length) + rand.Int63n(10000))
	for i := 0; i < length; i++ {
		rn := rand.Intn(len(char))
		w := char[rn : rn+1]
		t = append(t, w)
	}

	token := strings.Join(t, "")
	return token
}

func (req *DescribeTokenRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ValidateTokenRequest) Validate() error {
	return validate.Struct(req)
}

// NewDescribeTokenRequestByToken 查询详情请求
func NewDescribeTokenRequestByToken(token string) *DescribeTokenRequest {
	return &DescribeTokenRequest{
		DescribeBy:  DESCRIBE_BY_ACCESS_TOKEN,
		AccessToken: token,
	}
}

func NewDefaultToken() *Token {
	return &Token{}
}

func NewQueryTargetRequestFromHTTP(r *http.Request, name string) string {
	qs := r.URL.Query()
	return qs.Get(name)
}

func NewValidateTokenRequestFromHTTP(r *http.Request) *ValidateTokenRequest {
	qs := r.URL.Query()
	token := qs.Get("access_token")

	return &ValidateTokenRequest{
		AccessToken: token,
	}
}

func GetAccessTokenFromHTTP(r *http.Request) string {
	auth := r.Header.Get(ACCESS_TOKEN_HEADER_KEY)
	info := strings.Split(auth, " ")
	if len(info) > 1 {
		return info[1]
	}

	ck, err := r.Cookie(ACCESS_TOKEN_COOKIE_KEY)
	if err != nil {
		zap.L().Warnf("get tk from cookie: %s error, %s", ACCESS_TOKEN_COOKIE_KEY, err)
		return ""
	}

	return ck.Value
}

func NewValidateTokenRequest(token string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: token,
	}
}
