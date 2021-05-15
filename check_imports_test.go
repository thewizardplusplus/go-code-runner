package coderunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

func TestCheckImports(test *testing.T) {
	type args struct {
		code           string
		allowedImports []string
	}

	for _, data := range []struct {
		name      string
		args      args
		wantedErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				code: `
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
				`,
				allowedImports: []string{"fmt", "log"},
			},
			wantedErr: assert.NoError,
		},
		{
			name: "error with code parsing",
			args: args{
				code:           "incorrect",
				allowedImports: []string{"fmt", "log"},
			},
			wantedErr: assert.Error,
		},
		{
			name: "error with a disallowed import",
			args: args{
				code: `
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
				`,
				allowedImports: []string{"fmt"},
			},
			wantedErr: assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			pathToCode, err := systemutils.SaveTemporaryText(data.args.code)
			require.NoError(test, err)

			receivedErr := CheckImports(pathToCode, data.args.allowedImports)

			data.wantedErr(test, receivedErr)
		})
	}
}
