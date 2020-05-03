package keyregexdata

import "encoding/json"

// Expected what the test cases expect
type Expected struct {
	Output  string
	WantErr bool
}

// TODO: check param error vs system error

// TestCase how the test cases are structured
type TestCase struct {
	Input       string
	Expected    Expected
	Description string
}

type tests []TestCase

// TestCases a list of test cases for the conversion
var TestCases = tests{
	{
		"A-Z",
		Expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZ]]",
			false,
		},
		"simple upper case letter range",
	},
	{
		"[A-Z]",
		Expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZ]]",
			false,
		},
		"simple upper case letter range with square brackets",
	},
	{
		"a-z",
		Expected{
			"[[abcdefghijklmnopqrstuvwxyz]]",
			false,
		},
		"simple lower case letter range",
	},
	{
		"b-f",
		Expected{
			"[[bcdef]]",
			true,
		},
		"partial lower case letter range",
	},
	{
		"[a-z]",
		Expected{
			"[[abcdefghijklmnopqrstuvwxyz]]",
			false,
		},
		"simple lower case letter range with square brackets",
	},
	{
		"0-9",
		Expected{
			"[[0123456789]]",
			false,
		},
		"simple digits range",
	},
	{
		"5-8",
		Expected{
			"[[0123456789]]",
			true,
		},
		"partial digits range",
	},
	{
		"[0-9]",
		Expected{
			"[[0123456789]]",
			false,
		},
		"simple digits range with square brackets",
	},
	{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
		Expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789]]",
			false,
		},
		"input without regex",
	},
	{
		"[ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789]",
		Expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789]]",
			false,
		},
		"input without regex, with square brackets",
	},
	{
		"[ABCD][abcd]",
		Expected{
			"[[ABCD][abcd]]",
			false,
		},
		"double square brackets without regex",
	},
	{
		"[A-Z][a-z][0-9]",
		Expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZ][abcdefghijklmnopqrstuvwxyz][0123456789]]",
			false,
		},
		"trible square brackets with regex",
	},
	{
		"[0-5a-y][A-Dafg]",
		Expected{
			"[[012345abcdefghijklmnopqrstuvwxy][ABCDafg]]",
			false,
		},
		"complex example with partial regex, combined with no regex and double square brackets",
	},
	{
		"[%&/()=?]",
		Expected{
			"[[%&/()=?]]",
			false,
		},
		"special characters without regex",
	},
}

// FilePath defines where to export the testcases
func (t tests) FilePath() string {
	return "test/data/keyregexdata/export.json"
}

// FilePath defines where to export the testcases
func (t tests) Export() (out string, err error) {
	outByte, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return
	}
	return string(outByte), nil
}
