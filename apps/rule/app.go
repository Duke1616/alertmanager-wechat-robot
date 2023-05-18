package rule

import (
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "rule"
)

var (
	validate = validator.New()
)
