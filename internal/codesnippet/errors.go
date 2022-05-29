package codesnippet

import "errors"

var (
	ErrInsufficientHeader       = errors.New("Insufficient Header Information")
	ErrEmptyBody                = errors.New("Body is empty")
	ErrUnsupportedFileExtension = errors.New("Unsupported file extension")
)
