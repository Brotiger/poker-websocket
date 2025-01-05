package error

import "errors"

var ErrConnectTokenNotFound error
var ErrInvalidConnectToken error

func init() {
	ErrConnectTokenNotFound = errors.New("connect token not found")
	ErrInvalidConnectToken = errors.New("invalid connect token")
}
