package testrunner

import (
	"fmt"

	"github.com/pkg/errors"
	coderunner "github.com/thewizardplusplus/go-code-runner"
)

// TestCase ...
type TestCase struct {
	Input          string
	ExpectedOutput string
}

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
		"unexpected output: expected - %q, actual - %q",
		err.ExpectedOutput,
		err.ActualOutput,
	)
}

// RunCode ...
func RunCode(pathToExecutable string, testCases []TestCase) error {
	for _, testCase := range testCases {
		actualOutput, err := coderunner.RunCode(pathToExecutable, testCase.Input)
		if err != nil {
			return errors.Wrapf(
				err,
				"unable to run the test case (input - %q)",
				testCase.Input,
			)
		}
		if actualOutput != testCase.ExpectedOutput {
			return ErrUnexpectedOutput{TestCase: testCase, ActualOutput: actualOutput}
		}
	}

	return nil
}
