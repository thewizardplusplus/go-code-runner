package systemutils

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

// SaveTemporaryCode ...
func SaveTemporaryCode(code string) (path string, err error) {
	tempDir, err := ioutil.TempDir("", "code")
	if err != nil {
		return "", errors.Wrap(err, "unable to create a temporary directory")
	}

	tempFile := filepath.Join(tempDir, "code.go")
	if err := ioutil.WriteFile(tempFile, []byte(code), 0600); err != nil {
		return "", errors.Wrap(err, "unable to write the code")
	}

	return tempFile, nil
}
