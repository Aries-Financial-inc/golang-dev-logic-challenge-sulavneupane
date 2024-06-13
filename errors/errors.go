package errors

import (
	_ "errors"

	"github.com/pkg/errors"
)

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func New(message string) error {
	return errors.New(message)
}
