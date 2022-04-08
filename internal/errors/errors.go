package errors

import "github.com/pkg/errors"

// New returns a new error which indicates that the operation failed.
func New(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
