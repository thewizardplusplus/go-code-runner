package coderunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		// TODO: Add test cases.
	} {
		test.Run(data.name, func(test *testing.T) {
			pathToCode, err := SaveTemporaryCode(data.args.code)
			require.NoError(test, err)

			receivedErr := CheckImports(pathToCode, data.args.allowedImports)

			data.wantedErr(test, receivedErr)
		})
	}
}
