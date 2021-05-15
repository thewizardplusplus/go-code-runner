package systemutils

import (
	"context"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// RunCommand ...
func RunCommand(
	ctx context.Context,
	input string,
	command string,
	arguments ...string,
) (output string, err error) {
	commandInstance := exec.CommandContext(ctx, command, arguments...)
	commandInstance.Stdin = strings.NewReader(input)

	outputBytes, err := commandInstance.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			err = errors.Wrapf(err, "%q", exitErr.Stderr)
		}

		return "", errors.Wrap(err, "unable to run the command")
	}

	return string(outputBytes), nil
}
