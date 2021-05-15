package coderunner

import (
	"debug/elf"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

func TestCompileCode(test *testing.T) {
	const code = `package main; func main() { fmt.Println("Hello, World!") }`
	pathToCode, err := systemutils.SaveTemporaryCode(code)
	require.NoError(test, err)

	pathToExecutable, err := CompileCode(pathToCode, nil)
	require.NoError(test, err)

	codeContent, err := ioutil.ReadFile(pathToCode)
	require.NoError(test, err)

	// we do not use filepath.Split() because it leaves the separator
	dir, file := filepath.Dir(pathToExecutable), filepath.Base(pathToExecutable)
	assert.Equal(test, os.TempDir(), filepath.Dir(dir))
	assert.Regexp(test, `code\d+`, filepath.Base(dir))
	assert.Equal(test, "code", file)

	const wantedCodeContent = "package main\n\n" +
		"import \"fmt\"\n\n" +
		"func main() { fmt.Println(\"Hello, World!\") }\n"
	assert.Equal(test, wantedCodeContent, string(codeContent))

	_, err = elf.Open(pathToExecutable)
	assert.NoError(test, err)
}

func TestCompileCode_withDisallowedImport(test *testing.T) {
	const code = `package main; func main() { fmt.Println("Hello, World!") }`
	pathToCode, err := systemutils.SaveTemporaryCode(code)
	require.NoError(test, err)

	pathToExecutable, compileErr := CompileCode(pathToCode, []string{"log"})

	codeContent, err := ioutil.ReadFile(pathToCode)
	require.NoError(test, err)

	const wantedCodeContent = "package main\n\n" +
		"import \"fmt\"\n\n" +
		"func main() { fmt.Println(\"Hello, World!\") }\n"
	assert.Equal(test, wantedCodeContent, string(codeContent))

	const wantedCompileErrMessage = "failed import checking: " +
		`disallowed import "fmt"`
	assert.Empty(test, pathToExecutable)
	assert.EqualError(test, compileErr, wantedCompileErrMessage)
}
