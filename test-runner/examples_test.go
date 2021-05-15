package testrunner_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
	testrunner "github.com/thewizardplusplus/go-code-runner/test-runner"
)

func ExampleRunTestCases_success() {
	const code = `
		package main

		func main() {
			var x, y int
			fmt.Scan(&x, &y)

			fmt.Println(x + y)
		}
	`

	pathToCode, err := systemutils.SaveTemporaryText(code, ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	pathToExecutable, err := coderunner.CompileCode(pathToCode, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = testrunner.RunTestCases(ctx, pathToExecutable, []testrunner.TestCase{
		{Input: "5 12", ExpectedOutput: "17\n"},
		{Input: "23 42", ExpectedOutput: "65\n"},
	})
	fmt.Printf("%v\n", err)

	// Output:
	// <nil>
}

func ExampleRunTestCases_error() {
	const code = `
		package main

		func main() {
			var x, y int
			fmt.Scan(&x, &y)

			fmt.Println(x + y)
		}
	`

	pathToCode, err := systemutils.SaveTemporaryText(code, ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	pathToExecutable, err := coderunner.CompileCode(pathToCode, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = testrunner.RunTestCases(ctx, pathToExecutable, []testrunner.TestCase{
		{Input: "5 12", ExpectedOutput: "17\n"},
		{Input: "23 42", ExpectedOutput: "100\n"},
	})
	fmt.Printf("%v\n", err)

	// Output:
	// unexpected output (input - "23 42"): expected - "100\n", actual - "65\n"
}
