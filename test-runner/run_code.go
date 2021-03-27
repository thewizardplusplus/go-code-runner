package testrunner

import (
	"fmt"
)

// TestCase ...
type TestCase struct {
	Input          string
	ExpectedOutput string
}

// ErrTestCase ...
type ErrTestCase struct {
	TestCase

	ActualOutput string
}

// Error ...
func (err ErrTestCase) Error() string {
	return fmt.Sprintf(
		"unexpected output: expected - %q, actual - %q",
		err.ExpectedOutput,
		err.ActualOutput,
	)
}
