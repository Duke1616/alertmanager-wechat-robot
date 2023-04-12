package token

import (
	"github.com/go-playground/validator/v10"
	"math/rand"
	"strings"
	"time"
)

const (
	AppName = "token"
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
