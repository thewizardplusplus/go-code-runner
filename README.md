# go-code-runner

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-code-runner?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-code-runner)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-code-runner)](https://goreportcard.com/report/github.com/thewizardplusplus/go-code-runner)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-code-runner.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-code-runner)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-code-runner/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-code-runner)

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

`coderunner.CheckImports` (success):

```go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
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

	pathToCode, err := coderunner.SaveTemporaryCode(code)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	err = coderunner.CheckImports(pathToCode, []string{"fmt", "log"})
	fmt.Printf("%v\n", err)

	// Output:
	// <nil>
}
```

`coderunner.CheckImports` (error):

```go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
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

	pathToCode, err := coderunner.SaveTemporaryCode(code)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	err = coderunner.CheckImports(pathToCode, []string{"fmt"})
	fmt.Printf("%v\n", err)

	// Output:
	// disallowed import "log"
}
```

`coderunner.RunCode`:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
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

	pathToCode, err := coderunner.SaveTemporaryCode(code)
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
```

`testrunner.RunCode` (success):

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
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

	pathToCode, err := coderunner.SaveTemporaryCode(code)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	pathToExecutable, err := coderunner.CompileCode(pathToCode, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = testrunner.RunCode(ctx, pathToExecutable, []testrunner.TestCase{
		{Input: "5 12", ExpectedOutput: "17\n"},
		{Input: "23 42", ExpectedOutput: "65\n"},
	})
	fmt.Printf("%v\n", err)

	// Output:
	// <nil>
}
```

`testrunner.RunCode` (error):

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	coderunner "github.com/thewizardplusplus/go-code-runner"
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

	pathToCode, err := coderunner.SaveTemporaryCode(code)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(pathToCode)) // nolint: errcheck

	pathToExecutable, err := coderunner.CompileCode(pathToCode, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = testrunner.RunCode(ctx, pathToExecutable, []testrunner.TestCase{
		{Input: "5 12", ExpectedOutput: "17\n"},
		{Input: "23 42", ExpectedOutput: "100\n"},
	})
	fmt.Printf("%v\n", err)

	// Output:
	// unexpected output (input - "23 42"): expected - "100\n", actual - "65\n"
}
```

## License

The MIT License (MIT)

Copyright &copy; 2021 thewizardplusplus
