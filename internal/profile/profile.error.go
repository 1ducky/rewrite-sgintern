package profile

import "errors"

var (
	ErrNotFound        = errors.New("user not found")
	ErrDuplicateTag    = errors.New("tag already exists")
	ErrInvalidUsername = errors.New("invalid username")
)
