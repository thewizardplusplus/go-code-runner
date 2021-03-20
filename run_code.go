package coderunner

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// RunCode ...
func RunCode(pathToExecutable string, input string) (output string, err error) {
	cmd := exec.Command(pathToExecutable)
	cmd.Stdin = strings.NewReader(input)

	outputBytes, err := cmd.Output()
	if err != nil {
		err = wrapExitError(err)
		return "", errors.Wrap(err, "unable to run the code")
	}

	return string(outputBytes), nil
}
