package testrunner

import (
	"context"

	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

// TestCase ...
type TestCase struct {
	Input          string
	ExpectedOutput string
}

// RunTestCases ...
func RunTestCases(
	ctx context.Context,
	pathToExecutable string,
	testCases []TestCase,
) error {
	for _, testCase := range testCases {
		actualOutput, err :=
			systemutils.RunCommand(ctx, testCase.Input, pathToExecutable)
		if err != nil {
			return ErrFailedRunning{TestCase: testCase, ErrMessage: err.Error()}
		}
		if actualOutput != testCase.ExpectedOutput {
			return ErrUnexpectedOutput{TestCase: testCase, ActualOutput: actualOutput}
		}
	}

	return nil
}
