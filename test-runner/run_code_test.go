package testrunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrTestCase_Error(test *testing.T) {
	err := ErrTestCase{
		TestCase:     TestCase{Input: "input", ExpectedOutput: "expected output"},
		ActualOutput: "actual output",
	}

	const wantedErrMessage = "unexpected output: " +
		`expected - "expected output", actual - "actual output"`
	assert.EqualError(test, err, wantedErrMessage)
}
