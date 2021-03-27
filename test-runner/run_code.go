package testrunner

import (
	coderunner "github.com/thewizardplusplus/go-code-runner"
)

// TestCase ...
type TestCase struct {
	Input          string
	ExpectedOutput string
}

// RunCode ...
func RunCode(pathToExecutable string, testCases []TestCase) error {
	for _, testCase := range testCases {
		actualOutput, err := coderunner.RunCode(pathToExecutable, testCase.Input)
		if err != nil {
			return ErrFailedRunning{TestCase: testCase, ErrMessage: err.Error()}
		}
		if actualOutput != testCase.ExpectedOutput {
			return ErrUnexpectedOutput{TestCase: testCase, ActualOutput: actualOutput}
		}
	}

	return nil
}
