package coderunner_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
)

func ExampleRunCode() {
	const code = `
		package main

		func main() {
			var x, y int
			fmt.Scan(&x, &y)

			fmt.Println(x + y)
		}
	`

	pathToCode, err := coderunner.SaveTemporaryCode(code)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	pathToExecutable, err := coderunner.CompileCode(pathToCode)
	if err != nil {
		log.Fatal(err)
	}

	output, err := coderunner.RunCode(pathToExecutable, "2 3")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q\n", output)

	// Output:
	// "5\n"
}
