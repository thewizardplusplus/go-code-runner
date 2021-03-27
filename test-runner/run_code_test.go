package testrunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	coderunner "github.com/thewizardplusplus/go-code-runner"
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

func TestRunCode(test *testing.T) {
	type args struct {
		code      string
		testCases []TestCase
	}

	for _, data := range []struct {
		name      string
		args      args
		wantedErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	} {
		test.Run(data.name, func(test *testing.T) {
			pathToCode, err := coderunner.SaveTemporaryCode(data.args.code)
			require.NoError(test, err)

			pathToExecutable, err := coderunner.CompileCode(pathToCode)
			require.NoError(test, err)

			receivedErr := RunCode(pathToExecutable, data.args.testCases)
			require.NoError(test, err)

			data.wantedErr(test, receivedErr)
		})
	}
}
