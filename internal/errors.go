package internal

import "errors"

var (
	ErrEmptyParams = errors.New("empty params")
	ErrBadRequest  = errors.New("bad request")
	ErrInternal    = errors.New("internal server error")
	ErrNotFound    = errors.New("not found")
)
