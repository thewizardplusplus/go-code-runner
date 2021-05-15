package coderunner

import (
	"context"
	"os/exec"
	"path/filepath"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/pkg/errors"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

// CompileCode ...
func CompileCode(
	ctx context.Context,
	pathToCode string,
	allowedImports mapset.Set,
) (pathToExecutable string, err error) {
	if _, err = exec.LookPath("goimports"); err == nil {
		_, err = systemutils.RunCommand(ctx, "", "goimports", "-w", pathToCode)
		if err != nil {
			return "", errors.Wrap(err, "unable to prepare the code")
		}
	}

	if allowedImports != nil {
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
