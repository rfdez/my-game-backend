package errors

import "github.com/pkg/errors"

// New returns a new error which indicates that the operation failed.
func New(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
func Wrap(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}
