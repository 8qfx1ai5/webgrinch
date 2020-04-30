package encodetext

import (
	"testing"
)

type input struct {
	content string
	keyFrom string
	keyTo   string
}

type expected struct {
	output  string
	wantErr bool
}

type testCase struct {
	input       input
	expected    expected
	description string
}

/**
 * a list of test cases for the conversion
 */
var testCases = []testCase{
	{
		input{
			"Hello World!\n\nThis is an basic example to convert a plain text or message.",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		},
		expected{
			"",
			false,
		},
		"simple check",
	},
	{
		input{
			"",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		},
		expected{
			"",
			false,
		},
		"test empty input",
	},

	// TODO: test xml
	// TODO: test big input
	// TODO: test special characters like chineese
	// TODO: test other keys
	// TODO: test reverse translation
}

/**
 * validate functionality of the encoder
 */
func _TestEncodeText(t *testing.T) {

	for _, currentTest := range testCases {
		func(tc testCase) {
			t.Run(tc.description, func(t *testing.T) {
				observed, err := Run(tc.input.content, tc.input.keyFrom, tc.input.keyTo)
				if err != nil {
					if tc.expected.wantErr {
						return
					}
					t.Error(err)
				}

				// compare
				if observed != tc.expected.output {
					t.Errorf("\ninput='%s',\nkeyFrom='%s',\nkeyTo='%s',\nobserved='%s',\nexpected='%s',\ndescription: %s", tc.input.content, tc.input.keyFrom, tc.input.keyTo, observed, tc.expected.output, tc.description)
				}
			})
		}(currentTest)
	}
}

/**
 * check performance of the encoder
 */
func _BenchmarkEncoding(b *testing.B) {
	for x := 0; x < b.N; x++ {
		for _, testCase := range testCases {
			Run(testCase.input.content, testCase.input.keyFrom, testCase.input.keyTo)
		}
	}
}
