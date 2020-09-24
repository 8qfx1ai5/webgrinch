package enkey_test

import (
	"testing"

	"github.com/8qfx1ai5/webgrinch/internal/types/enkey"
	"github.com/8qfx1ai5/webgrinch/test/data/keyregexdata"
)

/**
 * validate functionality of the encoder
 */
func TestUseRegex(t *testing.T) {

	for _, currentTest := range keyregexdata.TestCases {
		func(tc keyregexdata.TestCase) {
			t.Run(tc.Description, func(t *testing.T) {
				var key = enkey.Key{}
				ok, err := key.UseRegex(tc.Input)
				if err != nil {
					if !ok && tc.Expected.WantErr {
						return
					}
					t.Error(err)
				}
				observed := key.GetSubsets()
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
