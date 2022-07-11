package globalErrors

import "errors"

var (
	ErrRecordNotFound           = errors.New("record not found")
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")
)
