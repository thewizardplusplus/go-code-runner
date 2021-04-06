package coderunner

import (
	"context"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// RunCode ...
func RunCode(
	ctx context.Context,
	pathToExecutable string,
	input string,
) (output string, err error) {
	cmd := exec.CommandContext(ctx, pathToExecutable) // nolint: gosec
	cmd.Stdin = strings.NewReader(input)

	outputBytes, err := cmd.Output()
	if err != nil {
		err = wrapExitError(err)
		return "", errors.Wrap(err, "unable to run the code")
	}

	return string(outputBytes), nil
}
