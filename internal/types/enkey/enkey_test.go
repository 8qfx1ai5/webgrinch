package enkey

import (
	"testing"
)

type expected struct {
	output  string
	wantErr bool
}

type testCase struct {
	input       string
	expected    expected
	description string
}

/**
 * a list of test cases for the conversion
 */
var testCasesKeyRegex = []testCase{
	{
		"A-Z",
		expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZ]]",
			false,
		},
		"simple upper case letter range",
	},
	{
		"[A-Z]",
		expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZ]]",
			false,
		},
		"simple upper case letter range with square brackets",
	},
	{
		"a-z",
		expected{
			"[[abcdefghijklmnopqrstuvwxyz]]",
			false,
		},
		"simple lower case letter range",
	},
	{
		"b-f",
		expected{
			"[[bcdef]]",
			true,
		},
		"partial lower case letter range",
	},
	{
		"[a-z]",
		expected{
			"[[abcdefghijklmnopqrstuvwxyz]]",
			false,
		},
		"simple lower case letter range with square brackets",
	},
	{
		"0-9",
		expected{
			"[[0123456789]]",
			false,
		},
		"simple digits range",
	},
	{
		"5-8",
		expected{
			"[[0123456789]]",
			true,
		},
		"partial digits range",
	},
	{
		"[0-9]",
		expected{
			"[[0123456789]]",
			false,
		},
		"simple digits range with square brackets",
	},
	{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
		expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789]]",
			false,
		},
		"input without regex",
	},
	{
		"[ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789]",
		expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789]]",
			false,
		},
		"input without regex, with square brackets",
	},
	{
		"[ABCD][abcd]",
		expected{
			"[[ABCD][abcd]]",
			false,
		},
		"double square brackets without regex",
	},
	{
		"[A-Z][a-z][0-9]",
		expected{
			"[[ABCDEFGHIJKLMNOPQRSTUVWXYZ][abcdefghijklmnopqrstuvwxyz][0123456789]]",
			false,
		},
		"trible square brackets with regex",
	},
	{
		"[0-5a-y][A-Dafg]",
		expected{
			"[[012345abcdefghijklmnopqrstuvwxy][ABCDafg]]",
			false,
		},
		"complex example with partial regex, combined with no regex and double square brackets",
	},
	{
		"[%&/()=?]",
		expected{
			"[[%&/()=?]]",
			false,
		},
		"special characters without regex",
	},
}

/**
 * validate functionality of the encoder
 */
func TestKeyRegexConversion(t *testing.T) {

	for _, currentTest := range testCasesKeyRegex {
		func(tc testCase) {
			t.Run(tc.description, func(t *testing.T) {
				observed, err := createSubsetsFromRegex(tc.input)
				if err != nil {
					if tc.expected.wantErr {
						return
					}
					t.Error(err)
				}
				observedString := convertSubsetToString(observed)
				if observedString != tc.expected.output {
					t.Errorf("\ninput=   '%s',\nobserved='%s',\nexpected='%s',\ndescription: %s", tc.input, observedString, tc.expected.output, tc.description)
				}
			})
		}(currentTest)
	}
}

// create readable string for comparison in the test
func convertSubsetToString(input [][]rune) string {
	out := ""
	for _, subset := range input {
		out += "[" + string(subset) + "]"
	}
	return "[" + out + "]"
}
