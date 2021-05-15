package systemutils

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

// SaveTemporaryText ...
func SaveTemporaryText(text string) (path string, err error) {
	tempDir, err := ioutil.TempDir("", "text")
	if err != nil {
		return "", errors.Wrap(err, "unable to create a temporary directory")
	}

	tempFile := filepath.Join(tempDir, "text.go")
	if err := ioutil.WriteFile(tempFile, []byte(text), 0600); err != nil {
		return "", errors.Wrap(err, "unable to write the text")
	}

	return tempFile, nil
}
