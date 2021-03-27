package testrunner

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	coderunner "github.com/thewizardplusplus/go-code-runner"
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
		{
			name: "success",
			args: args{
				code: `
					package main

					func main() {
						var x, y int
						fmt.Scan(&x, &y)

						fmt.Println(x + y)
					}
				`,
				testCases: []TestCase{
					{Input: "5 12", ExpectedOutput: "17\n"},
					{Input: "23 42", ExpectedOutput: "65\n"},
				},
			},
			wantedErr: assert.NoError,
		},
		{
			name: "error with failed running",
			args: args{
				code: `
					package main

					func main() {
						panic("error")
					}
				`,
				testCases: []TestCase{
					{Input: "5 12", ExpectedOutput: "17\n"},
					{Input: "23 42", ExpectedOutput: "65\n"},
				},
			},
			wantedErr: func(
				test assert.TestingT,
				err error,
				msgAndArgs ...interface{},
			) bool {
				if !assert.IsType(test, ErrFailedRunning{}, err) {
					return false
				}

				wantedTestCase := TestCase{Input: "5 12", ExpectedOutput: "17\n"}
				if !assert.Equal(test, wantedTestCase, err.(ErrFailedRunning).TestCase) {
					return false
				}

				return assert.True(test, strings.HasPrefix(
					err.(ErrFailedRunning).ErrMessage,
					"unable to run the code",
				))
			},
		},
		{
			name: "error with an unexpected output",
			args: args{
				code: `
					package main

					func main() {
						var x, y int
						fmt.Scan(&x, &y)

						fmt.Println(x + y)
					}
				`,
				testCases: []TestCase{
					{Input: "5 12", ExpectedOutput: "17\n"},
					{Input: "23 42", ExpectedOutput: "100\n"},
				},
			},
			wantedErr: func(
				test assert.TestingT,
				err error,
				msgAndArgs ...interface{},
			) bool {
				wantedErr := ErrUnexpectedOutput{
					TestCase:     TestCase{Input: "23 42", ExpectedOutput: "100\n"},
					ActualOutput: "65\n",
				}
				return assert.Equal(test, wantedErr, err)
			},
		},
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
