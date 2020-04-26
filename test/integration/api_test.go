package api_test

import (
	"net/http"
	"testing"
)

const (
	api string = "http://localhost:80"
)

type input struct {
	route     string
	method    string
	content   string
	keyFrom   string
	keyTo     string
	configCSS string
}

type expected struct {
	body string
	code int
}

/**
 * a list of test cases for the conversion
 */
var testCases = []struct {
	input       input
	expected    expected
	description string
}{
	{
		input{
			"",
			"GET",
			"<!-- this is a comment -->\n		<p>Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
			"vc",
		},
		expected{
			"",
			404,
		},
		"check server is reachable",
	},
}

/**
 * validate functionality of the encoder
 */
func TestAPI(t *testing.T) {

	for _, testCase := range testCases {
		var res *http.Response
		var err error

		switch testCase.input.method {
		case "GET":
			res, err = http.Get(api + testCase.input.route)
		default:
			t.Errorf("\nTest implementation invalid, method='%s' not found", testCase.input.method)
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}
		if res == nil {
			t.Errorf("\nAPI response not found")
			continue
		}

		// compare
		if res.StatusCode != testCase.expected.code {
			t.Errorf("\nAPI response code invalid. result='%d' expected='%d'", res.StatusCode, testCase.expected.code)
		}
	}
}
