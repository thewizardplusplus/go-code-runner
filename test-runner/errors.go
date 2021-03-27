package testrunner

import (
	"fmt"
)

// ErrFailedRunning ...
type ErrFailedRunning struct {
	TestCase

	ErrMessage string
}

// Error ...
func (err ErrFailedRunning) Error() string {
	return fmt.Sprintf(
		"failed running (input - %q): %s",
		err.Input,
		err.ErrMessage,
	)
}

// ErrUnexpectedOutput ...
type ErrUnexpectedOutput struct {
	TestCase

	ActualOutput string
}

// Error ...
func (err ErrUnexpectedOutput) Error() string {
	return fmt.Sprintf(
		"unexpected output (input - %q): expected - %q, actual - %q",
		err.Input,
		err.ExpectedOutput,
		err.ActualOutput,
	)
}
