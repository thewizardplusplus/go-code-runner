package coderunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunCode(test *testing.T) {
	const code = `
		package main

		func main() {
			var x, y int
			fmt.Scan(&x, &y)

			fmt.Println(x + y)
		}
	`

	pathToCode, err := SaveTemporaryCode(code)
	require.NoError(test, err)

	pathToExecutable, err := CompileCode(pathToCode)
	require.NoError(test, err)

	output, err := RunCode(pathToExecutable, "2 3")
	require.NoError(test, err)

	assert.Equal(test, "5\n", output)
}
