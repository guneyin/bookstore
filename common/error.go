package common

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExist  = errors.New("already exist")
	ErrInvalidUserId = errors.New("invalid user id")
)
