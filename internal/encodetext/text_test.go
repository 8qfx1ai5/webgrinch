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
			"KhooX GXzoe!\n\nITRB RB yv xyBRP hHyqmoh cX PXvlhzc y moyRv chHc Xz qhBByLh.",
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
	{
		input{
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		},
		expected{
			"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
			false,
		},
		"test all key characters",
	},
	{
		input{
			"Hello World!\n\nThis is an basic example to convert a plain text or message.\nABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"",
			"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		},
		expected{
			"Hello World!\n\nThis is an basic example to convert a plain text or message.\nABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			false,
		},
		"test empty 'from key'",
	},
	{
		input{
			"Hello World!\n\nThis is an basic example to convert a plain text or message.\nABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"",
		},
		expected{
			"Hello World!\n\nThis is an basic example to convert a plain text or message.\nABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			false,
		},
		"test empty 'to key'",
	},
	{
		input{
			"世BCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxy界",
			"世BCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxy界",
			"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		},
		expected{
			"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
			false,
		},
		"test chinese letters in 'from key'",
	},
	{
		input{
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"世FMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJ界",
		},
		expected{
			"世FMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJ界",
			false,
		},
		"test chinese letters in 'to key'",
	},
	{
		input{
			"$&.!#()?=§\"^°´`{}[]%+-_'*äüö@€:;<>µ~",
			"$&.!#()?=§\"^°´`{}[]%+-_'*äüö@€:;<>µ~",
			"~$&.!#()?=§\"^°´`{}[]%+-_'*äüö@€:;<>µ",
		},
		expected{
			"~$&.!#()?=§\"^°´`{}[]%+-_'*äüö@€:;<>µ",
			false,
		},
		"test special characters'",
	},
	{
		input{
			"EFGHIJKLMNOPQRSTUVWX YZabcdefghijklmnopqr stuvwxyz界世",
			"EFGHIJKLMNOPQRSTUVWX YZabcdefghijklmnopqr stuvwxyz界世",
			"aaaaaaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaaaaaa aaaaaaaaaa",
		},
		expected{
			"aaaaaaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaaaaaa aaaaaaaaaa",
			false,
		},
		"test overwrite, duplicates in 'key to'",
	},
	{
		input{
			"aaaaaaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaaaaaa aaaaaaaaaa",
			"aaaaaaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaaaaaa aaaaaaaaaa",
			"EFGHIJKLMNOPQRSTUVWX YZabcdefghijklmnopqr stuvwxyz",
		},
		expected{
			"zzzzzzzzzzzzzzzzzzzz zzzzzzzzzzzzzzzzzzzz zzzzzzzzzz",
			false,
		},
		"test overwrite, duplicates in 'key from'",
	},
}

/**
 * validate functionality of the encoder
 */
func TestEncodeText(t *testing.T) {

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
					t.Errorf("\ninput=   '%s',\nkeyFrom= '%s',\nkeyTo=   '%s',\nobserved='%s',\nexpected='%s',\ndescription: %s", tc.input.content, tc.input.keyFrom, tc.input.keyTo, observed, tc.expected.output, tc.description)
				}
			})
		}(currentTest)
	}
}

/**
 * check performance of the encoder
 */
func BenchmarkEncoding(b *testing.B) {
	for x := 0; x < b.N; x++ {
		for _, testCase := range testCases {
			Run(testCase.input.content, testCase.input.keyFrom, testCase.input.keyTo)
		}
	}
}
