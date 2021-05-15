package testrunner

import (
	"context"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRunTestCases(test *testing.T) {
	type args struct {
		ctx            context.Context
		testCases      []TestCase
		testCaseRunner TestCaseRunnerInterface
	}

	for _, data := range []struct {
		name      string
		args      args
		wantedErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				testCases: []TestCase{
					{Input: "5 12", ExpectedOutput: "17\n"},
					{Input: "23 42", ExpectedOutput: "65\n"},
				},
				testCaseRunner: func() TestCaseRunnerInterface {
					testCaseRunner := new(MockTestCaseRunnerInterface)
					testCaseRunner.
						On("RunTestCase", context.Background(), "5 12").
						Return("17\n", nil)
					testCaseRunner.
						On("RunTestCase", context.Background(), "23 42").
						Return("65\n", nil)

					return testCaseRunner
				}(),
			},
			wantedErr: nil,
		},
		{
			name: "error with failed running",
			args: args{
				ctx: context.Background(),
				testCases: []TestCase{
					{Input: "5 12", ExpectedOutput: "17\n"},
					{Input: "23 42", ExpectedOutput: "65\n"},
				},
				testCaseRunner: func() TestCaseRunnerInterface {
					testCaseRunner := new(MockTestCaseRunnerInterface)
					testCaseRunner.
						On("RunTestCase", context.Background(), "5 12").
						Return("", iotest.ErrTimeout)

					return testCaseRunner
				}(),
			},
			wantedErr: ErrFailedRunning{
				TestCase:   TestCase{Input: "5 12", ExpectedOutput: "17\n"},
				ErrMessage: iotest.ErrTimeout.Error(),
			},
		},
		{
			name: "error with an unexpected output",
			args: args{
				ctx: context.Background(),
				testCases: []TestCase{
					{Input: "5 12", ExpectedOutput: "17\n"},
					{Input: "23 42", ExpectedOutput: "65\n"},
				},
				testCaseRunner: func() TestCaseRunnerInterface {
					testCaseRunner := new(MockTestCaseRunnerInterface)
					testCaseRunner.
						On("RunTestCase", context.Background(), "5 12").
						Return("100\n", nil)

					return testCaseRunner
				}(),
			},
			wantedErr: ErrUnexpectedOutput{
				TestCase:     TestCase{Input: "5 12", ExpectedOutput: "17\n"},
				ActualOutput: "100\n",
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			receivedErr := RunTestCases(
				data.args.ctx,
				data.args.testCases,
				data.args.testCaseRunner.RunTestCase,
			)

			mock.AssertExpectationsForObjects(test, data.args.testCaseRunner)
			assert.Equal(test, data.wantedErr, receivedErr)
		})
	}
}
