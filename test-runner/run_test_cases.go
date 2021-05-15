package testrunner

import (
	"context"
)

// TestCase ...
type TestCase struct {
	Input          string
	ExpectedOutput string
}

// TestCaseRunner ...
type TestCaseRunner func(ctx context.Context, input string) (
	output string,
	err error,
)

// RunTestCases ...
func RunTestCases(
	ctx context.Context,
	testCases []TestCase,
	testCaseRunner TestCaseRunner,
) error {
	for _, testCase := range testCases {
		actualOutput, err := testCaseRunner(ctx, testCase.Input)
		if err != nil {
			return ErrFailedRunning{TestCase: testCase, ErrMessage: err.Error()}
		}
		if actualOutput != testCase.ExpectedOutput {
			return ErrUnexpectedOutput{TestCase: testCase, ActualOutput: actualOutput}
		}
	}

	return nil
}
