package testrunner

import (
	"context"
)

//go:generate mockery --name=TestCaseRunnerInterface --inpackage --case=underscore --testonly

// TestCaseRunnerInterface ...
//
// It is used only for mock generating.
//
type TestCaseRunnerInterface interface {
	RunTestCase(ctx context.Context, input string) (output string, err error)
}
