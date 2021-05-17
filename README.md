# go-code-runner

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-code-runner?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-code-runner)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-code-runner)](https://goreportcard.com/report/github.com/thewizardplusplus/go-code-runner)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-code-runner.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-code-runner)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-code-runner/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-code-runner)

The library that implements the compiling and running of a code written in the Go programming language and running of a test case set for the compiled code (i.e. the executable file).

## Features

- saving of a code to a temporary file:
  - storing of the temporary file with the code to an individual temporary directory;
- checking of package imports used in the code written in the Go programming language:
  - checking based on the set of allowed imports;
- compiling of a code written in the Go programming language:
  - automatic importing of the packages used in the code (optionally);
  - checking of package imports used in the code (optionally):
    - checking based on the set of allowed imports;
  - enriching of an error of the external command running by an output from the stderr stream;
  - support for the running time management (via the Go context);
- running of the compiled code (i.e. the executable file):
  - passing of a custom input as the stdin stream;
  - returning of an output from the stdout stream;
  - enriching of an error of the external command running by an output from the stderr stream;
  - support for the running time management (via the Go context);
- running of a test case set for the compiled code (i.e. the executable file):
  - representation of a test case:
    - input;
    - expected output;
  - checking of an actual output in each test case:
    - returning of the [sentinel errors](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully#sentinel%20errors):
      - failed running &mdash; it returns on a running error;
      - unexpected output &mdash; it returns when the expected and actual outputs do not match;
  - support for the running time management (via the Go context).

## Installation

Prepare the directory:

```
$ mkdir --parents "$(go env GOPATH)/src/github.com/thewizardplusplus/"
$ cd "$(go env GOPATH)/src/github.com/thewizardplusplus/"
```

Clone this repository:

```
$ git clone https://github.com/thewizardplusplus/go-code-runner.git
$ cd go-code-runner
```

Install dependencies with the [dep](https://golang.github.io/dep/) tool:

```
$ dep ensure -vendor-only
```

## Example

`coderunner.CheckImports()` (success):

```go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set"
	coderunner "github.com/thewizardplusplus/go-code-runner"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

func main() {
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

	pathToCode, err := systemutils.SaveTemporaryText(code, ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	err = coderunner.CheckImports(pathToCode, mapset.NewSet("fmt", "log"))
	fmt.Printf("%v\n", err)

	// Output:
	// <nil>
}
```

`coderunner.CheckImports()` (error):

```go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set"
	coderunner "github.com/thewizardplusplus/go-code-runner"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

func main() {
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

	pathToCode, err := systemutils.SaveTemporaryText(code, ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	err = coderunner.CheckImports(pathToCode, mapset.NewSet("log"))
	fmt.Printf("%v\n", err)

	// Output:
	// disallowed import "fmt"
}
```

`coderunner.CompileCode()`:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

func main() {
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

	pathToExecutable, err :=
		coderunner.CompileCode(context.Background(), pathToCode, nil)
	if err != nil {
		log.Fatal(err)
	}

	output, err :=
		systemutils.RunCommand(context.Background(), "2 3", pathToExecutable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q\n", output)

	// Output:
	// "5\n"
}
```

`testrunner.RunTestCases()` (success):

```go
package main

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

func main() {
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

	pathToExecutable, err :=
		coderunner.CompileCode(context.Background(), pathToCode, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = testrunner.RunTestCases(
		context.Background(),
		[]testrunner.TestCase{
			{Input: "5 12", ExpectedOutput: "17\n"},
			{Input: "23 42", ExpectedOutput: "65\n"},
		},
		func(ctx context.Context, input string) (output string, err error) {
			return systemutils.RunCommand(ctx, input, pathToExecutable)
		},
	)
	fmt.Printf("%v\n", err)

	// Output:
	// <nil>
}
```

`testrunner.RunTestCases()` (error):

```go
package main

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

func main() {
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

	pathToExecutable, err :=
		coderunner.CompileCode(context.Background(), pathToCode, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = testrunner.RunTestCases(
		context.Background(),
		[]testrunner.TestCase{
			{Input: "5 12", ExpectedOutput: "17\n"},
			{Input: "23 42", ExpectedOutput: "100\n"},
		},
		func(ctx context.Context, input string) (output string, err error) {
			return systemutils.RunCommand(ctx, input, pathToExecutable)
		},
	)
	fmt.Printf("%v\n", err)

	// Output:
	// unexpected output (input - "23 42"): expected - "100\n", actual - "65\n"
}
```

## License

The MIT License (MIT)

Copyright &copy; 2021 thewizardplusplus
