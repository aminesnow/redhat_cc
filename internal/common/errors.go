package common

import (
	"github.com/pkg/errors"
)

// ErrInternalError is returned if an unexpected error is raised or relates to internal workings of the function.
type ErrInternalError struct{ error }

// ErrUnavailable is returned an entity was not found.
type ErrNotFoundError struct{ error }

/*
 * Constructors
 */
func NewErrInternalError(err error, msg string, args ...interface{}) ErrInternalError {
	return ErrInternalError{errors.Wrapf(err, msg, args...)}
}

func NewErrInternalErrorMsg(msg string, args ...interface{}) ErrInternalError {
	return ErrInternalError{errors.Errorf(msg, args...)}
}

func NewErrNotFoundError(name string, id string) ErrNotFoundError {
	return ErrNotFoundError{errors.Errorf("entity %s:%s was not found", name, id)}
}
