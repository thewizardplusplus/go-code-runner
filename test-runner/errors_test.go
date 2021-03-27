package testrunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrFailedRunning_Error(test *testing.T) {
	err := ErrFailedRunning{
		TestCase:   TestCase{Input: "input", ExpectedOutput: "expected output"},
		ErrMessage: "error",
	}

	const wantedErrMessage = `failed running (input - "input"): error`
	assert.EqualError(test, err, wantedErrMessage)
}

func TestErrUnexpectedOutput_Error(test *testing.T) {
	err := ErrUnexpectedOutput{
		TestCase:     TestCase{Input: "input", ExpectedOutput: "expected output"},
		ActualOutput: "actual output",
	}

	const wantedErrMessage = "unexpected output: " +
		`expected - "expected output", actual - "actual output"`
	assert.EqualError(test, err, wantedErrMessage)
}
