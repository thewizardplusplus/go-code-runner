package systemutils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(test *testing.T) {
	type args struct {
		ctx       context.Context
		input     string
		command   string
		arguments []string
	}

	for _, data := range []struct {
		name         string
		args         args
		wantedOutput string
		wantedErr    assert.ErrorAssertionFunc
	}{
		{
			name: "success without an input",
			args: args{
				ctx:       context.Background(),
				input:     "",
				command:   "echo",
				arguments: []string{"test"},
			},
			wantedOutput: "test\n",
			wantedErr:    assert.NoError,
		},
		{
			name: "success with an input",
			args: args{
				ctx:       context.Background(),
				input:     "test",
				command:   "cat",
				arguments: nil,
			},
			wantedOutput: "test",
			wantedErr:    assert.NoError,
		},
		{
			name: "error without a stderr message",
			args: args{
				ctx:       context.Background(),
				input:     "",
				command:   "non-existent",
				arguments: nil,
			},
			wantedOutput: "",
			wantedErr:    assert.Error,
		},
		{
			name: "error with an empty stderr message",
			args: args{
				ctx:       context.Background(),
				input:     "",
				command:   "false",
				arguments: nil,
			},
			wantedOutput: "",
			wantedErr:    assert.Error,
		},
		{
			name: "error with a non-empty stderr message",
			args: args{
				ctx:       context.Background(),
				input:     "",
				command:   "cat",
				arguments: []string{"non-existent"},
			},
			wantedOutput: "",
			wantedErr:    assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			receivedOutput, receivedErr := RunCommand(
				data.args.ctx,
				data.args.input,
				data.args.command,
				data.args.arguments...,
			)

			assert.Equal(test, data.wantedOutput, receivedOutput)
			data.wantedErr(test, receivedErr)
		})
	}
}
