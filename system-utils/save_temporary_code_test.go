package systemutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveTemporaryCode(test *testing.T) {
	const code = "test"
	path, err := SaveTemporaryCode(code)
	require.NoError(test, err)

	content, err := ioutil.ReadFile(path)
	require.NoError(test, err)

	// we do not use filepath.Split() because it leaves the separator
	dir, file := filepath.Dir(path), filepath.Base(path)
	assert.Equal(test, os.TempDir(), filepath.Dir(dir))
	assert.Regexp(test, `code\d+`, filepath.Base(dir))
	assert.Equal(test, "code.go", file)
	assert.Equal(test, code, string(content))
}
