package errors

import (
	"errors"
	"fmt"
)

const (
	logOutputError      = "Unsupported log output:"
	notImplementedError = "Not yet implemented:"
)

var (
	ErrLogLevel = errors.New("Could not set configured log level")
)

func ErrLogOutput(output string) error {
	return fmt.Errorf("%s %s", logOutputError, output)
}

func ErrNotImplementedE(name string) error {
	return fmt.Errorf("%s %s", notImplementedError, name)
}
