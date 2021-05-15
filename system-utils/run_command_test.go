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
		// TODO: Add test cases.
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
