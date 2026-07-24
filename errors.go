package gedis

import (
	"errors"

	"github.com/anteiro255/gedis/pkg/protocol/status"
)

var (
	// nil is used for github.com/anteiro255/gedis/pkg/protocol/status.OK

	// ErrNoSuchKey is returned when the requested key does not exist.
	ErrNoSuchKey = errors.New("gedis: key does not exist") // github.com/anteiro255/gedis/pkg/protocol/status.NoSuchKey

	// ErrKeyAlreadyExists is returned when trying to create a key that is already set.
	ErrKeyAlreadyExists = errors.New("gedis: key already exists") // github.com/anteiro255/gedis/pkg/protocol/status.KeyAlreadyExists

	// ErrWrongInput is returned when arguments or payload formats are invalid.
	ErrWrongInput = errors.New("gedis: wrong input or invalid argument") // github.com/anteiro255/gedis/pkg/protocol/status.WrongInput

	// ErrInternalError is returned when the server encounters an unhandled internal error.
	ErrInternalError = errors.New("gedis: internal server error") // github.com/anteiro255/gedis/pkg/protocol/status.InternalError

	// ErrDeadlineExceeded is returned when the operation times out on the server side.
	ErrDeadlineExceeded = errors.New("gedis: deadline exceeded") // github.com/anteiro255/gedis/pkg/protocol/status.DeadlineExceeded

	// ErrUnknownStatus is returned when receiving an unrecognized status code.
	ErrUnknownStatus = errors.New("gedis: unknown status code") // github.com/anteiro255/gedis/pkg/protocol/status.UnknownStatus
)

func mapStatusToError(stat status.Status) error {
	switch stat {
	case status.OK:
		return nil
	case status.NoSuchKey:
		return ErrNoSuchKey
	case status.KeyAlreadyExists:
		return ErrKeyAlreadyExists
	case status.WrongInput:
		return ErrWrongInput
	case status.InternalError:
		return ErrInternalError
	case status.DeadlineExceeded:
		return ErrDeadlineExceeded
	default:
		return ErrUnknownStatus
	}
}
