package error

import "errors"

var ErrUserNotFound error

func init() {
	ErrUserNotFound = errors.New("user not found")
}
