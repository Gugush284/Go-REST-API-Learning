package Constants

import (
	"errors"
)

type ctxKey int8

var (
	ErrRecordNotFound           = errors.New("record not found")
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")
	ErrNotAuthenticated         = errors.New("not authenticated")
)

const SessionName = "activesession"

const (
	CtxKeyUser ctxKey = iota
	CtxKeyId
)
