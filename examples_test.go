package coderunner_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

func ExampleCheckImports_success() {
	const code = `
		package main

		import (
			"fmt"
			"log"
		)

		func main() {
			var x, y int
			if _, err := fmt.Scan(&x, &y); err != nil {
				log.Fatal(err)
			}

			fmt.Println(x + y)
		}
	`

	pathToCode, err := systemutils.SaveTemporaryText(code)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	err = coderunner.CheckImports(pathToCode, []string{"fmt", "log"})
	fmt.Printf("%v\n", err)

	// Output:
	// <nil>
}

func ExampleCheckImports_error() {
	const code = `
		package main

		import (
			"fmt"
			"log"
		)

		func main() {
			var x, y int
			if _, err := fmt.Scan(&x, &y); err != nil {
				log.Fatal(err)
			}

			fmt.Println(x + y)
		}
	`

	pathToCode, err := systemutils.SaveTemporaryText(code)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	err = coderunner.CheckImports(pathToCode, []string{"fmt"})
	fmt.Printf("%v\n", err)

	// Output:
	// disallowed import "log"
}

func ExampleRunCode() {
	const code = `
		package main

		func main() {
			var x, y int
			fmt.Scan(&x, &y)

			fmt.Println(x + y)
		}
	`

	pathToCode, err := systemutils.SaveTemporaryText(code)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	pathToExecutable, err := coderunner.CompileCode(pathToCode, nil)
	if err != nil {
		log.Fatal(err)
	}

	output, err :=
		coderunner.RunCode(context.Background(), pathToExecutable, "2 3")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q\n", output)

	// Output:
	// "5\n"
}
