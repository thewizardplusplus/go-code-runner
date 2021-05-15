package systemutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveTemporaryText(test *testing.T) {
	path, err := SaveTemporaryText("test", ".txt")
	require.NoError(test, err)

	content, err := ioutil.ReadFile(path)
	require.NoError(test, err)

	// we do not use filepath.Split() because it leaves the separator
	dir, file := filepath.Dir(path), filepath.Base(path)
	assert.Equal(test, os.TempDir(), filepath.Dir(dir))
	assert.Regexp(test, `text\d+`, filepath.Base(dir))
	assert.Equal(test, "text.txt", file)
	assert.Equal(test, "test", string(content))
}
