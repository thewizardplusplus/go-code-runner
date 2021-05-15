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
		// TODO: Add test cases.
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
