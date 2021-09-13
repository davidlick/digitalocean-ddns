package digitalocean

import "errors"

var (
	ErrNotFound    = errors.New("not found")
	ErrClientError = errors.New("client error")
	ErrServerError = errors.New("server error")
)
