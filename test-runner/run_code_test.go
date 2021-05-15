package testrunner

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	coderunner "github.com/thewizardplusplus/go-code-runner"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

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
					"unable to run the command",
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
			pathToCode, err := systemutils.SaveTemporaryText(data.args.code, ".go")
			require.NoError(test, err)

			pathToExecutable, err := coderunner.CompileCode(pathToCode, nil)
			require.NoError(test, err)

			receivedErr :=
				RunCode(context.Background(), pathToExecutable, data.args.testCases)
			require.NoError(test, err)

			data.wantedErr(test, receivedErr)
		})
	}
}

func TestRunCode_withTimeout(test *testing.T) {
	const code = `
		package main

		func main() {
			// sleep forever
			for {
				runtime.Gosched()
			}
		}
	`

	pathToCode, err := systemutils.SaveTemporaryText(code, ".go")
	require.NoError(test, err)

	pathToExecutable, err := coderunner.CompileCode(pathToCode, nil)
	require.NoError(test, err)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	receivedErr := RunCode(ctx, pathToExecutable, []TestCase{
		{Input: "5 12", ExpectedOutput: "17\n"},
		{Input: "23 42", ExpectedOutput: "65\n"},
	})

	assert.Error(test, receivedErr)
}
