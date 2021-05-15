package systemutils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// SaveTemporaryText ...
func SaveTemporaryText(text string, extension string) (path string, err error) {
	tempDir, err := ioutil.TempDir("", "text")
	if err != nil {
		return "", errors.Wrap(err, "unable to create a temporary directory")
	}

	tempFile := filepath.Join(tempDir, "text"+extension)
	file, err := os.OpenFile(tempFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return "", errors.Wrap(err, "unable to create a temporary file")
	}
	defer file.Close()

	if _, err := io.WriteString(file, text); err != nil {
		return "", errors.Wrap(err, "unable to write the text")
	}

	return tempFile, nil
}
