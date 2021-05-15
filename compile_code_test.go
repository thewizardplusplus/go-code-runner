package coderunner

import (
	"context"
	"io/ioutil"
	"testing"

	mapset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	systemutils "github.com/thewizardplusplus/go-code-runner/system-utils"
)

func TestCompileCode(test *testing.T) {
	type args struct {
		ctx            context.Context
		code           string
		allowedImports mapset.Set
		input          string
	}

	for _, data := range []struct {
		name             string
		args             args
		wantedErr        assert.ErrorAssertionFunc
		wantPreparedCode string
		wantOutput       string
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				code: `
					package main

					func main() {
						var x, y int
						if _, err := fmt.Scan(&x, &y); err != nil {
							log.Fatal(err)
						}

						fmt.Println(x + y)
					}
				`,
				allowedImports: nil,
				input:          "5 12",
			},
			wantPreparedCode: "package main\n" +
				"\n" +
				"import (\n" +
				"\t\"fmt\"\n" +
				"\t\"log\"\n" +
				")\n" +
				"\n" +
				"func main() {\n" +
				"\tvar x, y int\n" +
				"\tif _, err := fmt.Scan(&x, &y); err != nil {\n" +
				"\t\tlog.Fatal(err)\n" +
				"\t}\n" +
				"\n" +
				"\tfmt.Println(x + y)\n" +
				"}\n",
			wantOutput: "17\n",
			wantedErr:  assert.NoError,
		},
		{
			name: "error with code preparing",
			args: args{
				ctx: context.Background(),
				code: `
					package main

					func main() {
				`,
				allowedImports: nil,
				input:          "5 12",
			},
			wantPreparedCode: "",
			wantOutput:       "",
			wantedErr:        assert.Error,
		},
		{
			name: "error with import checking",
			args: args{
				ctx: context.Background(),
				code: `
					package main

					func main() {
						var x, y int
						if _, err := fmt.Scan(&x, &y); err != nil {
							log.Fatal(err)
						}

						fmt.Println(x + y)
					}
				`,
				allowedImports: mapset.NewSet("log"),
				input:          "5 12",
			},
			wantPreparedCode: "package main\n" +
				"\n" +
				"import (\n" +
				"\t\"fmt\"\n" +
				"\t\"log\"\n" +
				")\n" +
				"\n" +
				"func main() {\n" +
				"\tvar x, y int\n" +
				"\tif _, err := fmt.Scan(&x, &y); err != nil {\n" +
				"\t\tlog.Fatal(err)\n" +
				"\t}\n" +
				"\n" +
				"\tfmt.Println(x + y)\n" +
				"}\n",
			wantOutput: "",
			wantedErr:  assert.Error,
		},
		{
			name: "error with code compiling",
			args: args{
				ctx: context.Background(),
				code: `
					package main

					func main() {
						var x, y int
					}
				`,
				allowedImports: nil,
				input:          "5 12",
			},
			wantPreparedCode: "package main\n" +
				"\n" +
				"func main() {\n" +
				"\tvar x, y int\n" +
				"}\n",
			wantOutput: "",
			wantedErr:  assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			pathToCode, err := systemutils.SaveTemporaryText(data.args.code, ".go")
			require.NoError(test, err)

			pathToExecutable, receivedErr :=
				CompileCode(data.args.ctx, pathToCode, data.args.allowedImports)

			data.wantedErr(test, receivedErr)

			if data.wantPreparedCode != "" {
				preparedCode, err := ioutil.ReadFile(pathToCode)
				require.NoError(test, err)

				assert.Equal(test, data.wantPreparedCode, string(preparedCode))
			}

			if data.wantOutput != "" {
				output, err := systemutils.RunCommand(
					context.Background(),
					data.args.input,
					pathToExecutable,
				)
				require.NoError(test, err)

				assert.Equal(test, data.wantOutput, output)
			}
		})
	}
}
