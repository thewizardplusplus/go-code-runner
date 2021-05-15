package coderunner

import (
	"go/parser"
	"go/token"
	"strconv"

	mapset "github.com/deckarep/golang-set"
	"github.com/pkg/errors"
)

// CheckImports ...
func CheckImports(pathToCode string, allowedImports mapset.Set) error {
	fileSet := token.NewFileSet()
	ast, err := parser.ParseFile(fileSet, pathToCode, nil, parser.ImportsOnly)
	if err != nil {
		return errors.Wrap(err, "unable to parse the code")
	}

	for _, importSpec := range ast.Imports {
		importPath, err := strconv.Unquote(importSpec.Path.Value)
		if err != nil {
			return errors.Wrap(err, "unable to unquote the import path")
		}
		if !allowedImports.Contains(importPath) {
			return errors.Errorf("disallowed import %q", importPath)
		}
	}

	return nil
}
