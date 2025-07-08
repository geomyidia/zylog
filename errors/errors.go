// Package errors provides custom error types and functions for the application.
package errors

import (
	"errors"
	"fmt"
)

const (
	logOutputError      = "unsupported log output:"
	notImplementedError = "not yet implemented:"
)

// Errors
var (
	ErrLogLevel = errors.New("could not set configured log level")
)

// ErrLogOutput returns an error indicating that the specified log output is unsupported.
func ErrLogOutput(output string) error {
	return fmt.Errorf("%s %s", logOutputError, output)
}

// ErrNotImplemented returns an error indicating that a feature is not yet implemented.
func ErrNotImplemented(name string) error {
	return fmt.Errorf("%s %s", notImplementedError, name)
}
