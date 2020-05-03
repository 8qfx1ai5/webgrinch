package testcases

// TestCases is the basic structure of all test cases
type TestCases interface {
	Export() (string, error)
	FilePath() string
}
