package coderunner

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

// CompileCode ...
func CompileCode(
	pathToCode string,
	allowedImports []string,
) (pathToExecutable string, err error) {
	ctx := context.Background()
	_, err = systemutils.RunCommand(ctx, "", "goimports", "-w", pathToCode)
	if err != nil {
		return "", errors.Wrap(err, "unable to prepare the code")
	}

	if len(allowedImports) != 0 {
		if err = CheckImports(pathToCode, allowedImports); err != nil {
			return "", errors.Wrap(err, "failed import checking")
		}
	}

	pathToExecutable = strings.TrimSuffix(pathToCode, filepath.Ext(pathToCode))
	_, err = systemutils.RunCommand(
		ctx,
		"",
		"go",
		"build",
		"-o",
		pathToExecutable,
		pathToCode,
	)
	if err != nil {
		return "", errors.Wrap(err, "unable to compile the code")
	}

	return pathToExecutable, nil
}
