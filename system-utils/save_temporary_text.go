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
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(tempFile, flags, 0600) // nolint: gosec
	if err != nil {
		return "", errors.Wrap(err, "unable to create a temporary file")
	}
	defer file.Close() // nolint: errcheck, gosec

	if _, err := io.WriteString(file, text); err != nil {
		return "", errors.Wrap(err, "unable to write the text")
	}

	return tempFile, nil
}
