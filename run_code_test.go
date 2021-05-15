package coderunner

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
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

	pathToCode, err := systemutils.SaveTemporaryText(code, ".go")
	require.NoError(test, err)

	pathToExecutable, err := CompileCode(pathToCode, nil)
	require.NoError(test, err)

	output, err := RunCode(context.Background(), pathToExecutable, "2 3")
	require.NoError(test, err)

	assert.Equal(test, "5\n", output)
}

func TestRunCode_withTimeout(test *testing.T) {
	const code = `
		package main

		func main() {
			// sleep forever
			for {
				runtime.Gosched()
			}
		}
	`

	pathToCode, err := systemutils.SaveTemporaryText(code, ".go")
	require.NoError(test, err)

	pathToExecutable, err := CompileCode(pathToCode, nil)
	require.NoError(test, err)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	output, err := RunCode(ctx, pathToExecutable, "2 3")

	assert.Equal(test, "", output)
	assert.Error(test, err)
}
