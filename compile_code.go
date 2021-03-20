package coderunner

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// CompileCode ...
func CompileCode(pathToCode string) (pathToExecutable string, err error) {
	if _, err := exec.Command("goimports", "-w", pathToCode).Output(); err != nil {
		return "", errors.Wrap(err, "unable to prepare the code")
	}

	if _, err := exec.Command("go", "build", pathToCode).Output(); err != nil {
		return "", errors.Wrap(err, "unable to compile the code")
	}

	pathToExecutable = strings.TrimSuffix(pathToCode, filepath.Ext(pathToCode))
	return pathToExecutable, nil
}
