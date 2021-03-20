package coderunner

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// CompileCode ...
func CompileCode(pathToCode string) (pathToExecutable string, err error) {
	_, err = exec.Command("goimports", "-w", pathToCode).Output() // nolint: gosec
	if err != nil {
		err = wrapExitError(err)
		return "", errors.Wrap(err, "unable to prepare the code")
	}

	pathToExecutable = strings.TrimSuffix(pathToCode, filepath.Ext(pathToCode))
	_, err = exec. // nolint: gosec
			Command("go", "build", "-o", pathToExecutable, pathToCode).
			Output()
	if err != nil {
		err = wrapExitError(err)
		return "", errors.Wrap(err, "unable to compile the code")
	}

	return pathToExecutable, nil
}

func wrapExitError(err error) error {
	if exitErr, ok := err.(*exec.ExitError); ok {
		err = errors.Wrapf(err, "%q", string(exitErr.Stderr))
	}

	return err
}
