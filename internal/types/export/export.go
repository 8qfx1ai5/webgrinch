package export

// Export is the basic structure of all test cases
type Export interface {
	Export() (string, error)
	FileName() string
}
