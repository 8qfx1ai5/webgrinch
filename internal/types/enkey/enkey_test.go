package enkey

import (
	"testing"

	"github.com/8qfx1ai5/viewcrypt/test/data/keyregexdata"
)

/**
 * validate functionality of the encoder
 */
func TestKeyRegexConversion(t *testing.T) {

	for _, currentTest := range keyregexdata.TestCasesKeyRegex {
		func(tc keyregexdata.TestCase) {
			t.Run(tc.Description, func(t *testing.T) {
				observed, err := createSubsetsFromRegex(tc.Input)
				if err != nil {
					if tc.Expected.WantErr {
						return
					}
					t.Error(err)
				}
				observedString := convertSubsetToString(observed)
				if observedString != tc.Expected.Output {
					t.Errorf("\ninput=   '%s',\nobserved='%s',\nexpected='%s',\ndescription: %s", tc.Input, observedString, tc.Expected.Output, tc.Description)
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
