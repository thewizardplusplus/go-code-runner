package testrunner

// TestCase ...
type TestCase struct {
	Input          string
	ExpectedOutput string
}

// ErrTestCase ...
type ErrTestCase struct {
	TestCase

	ActualOutput string
}
