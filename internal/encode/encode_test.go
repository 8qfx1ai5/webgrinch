package encode

import (
	"strings"
	"testing"
)

type translations []struct {
	original    string
	replacement string
}

var testCases = []struct {
	input        string
	keyFrom      string
	keyTo        string
	configCSS    string
	translations translations
	description  string
}{
	{
		"<!-- this is a comment -->\n		<p>Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		"vc",
		translations{
			{"<p>Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>", "<p class=\"vc\">jXzhq RmBrq eXoXz BRc yqhc, PXvBhPchcrz <b>yeRmRBRPRvL</b> hoRc. fhmhooyc, ehohvRcR!</p>"},
		},
		"simple check for single <p> + comment",
	},
	{
		"<!-- this is a comment -->\n		<p>Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		"",
		translations{
			{"<p>Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>", "<p>jXzhq RmBrq eXoXz BRc yqhc, PXvBhPchcrz <b>yeRmRBRPRvL</b> hoRc. fhmhooyc, ehohvRcR!</p>"},
		},
		"do not add class attribute if class is empty",
	},
	{
		"<!-- this is a comment -->\n		<p class=\"foo\">Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		"vc",
		translations{
			{"<p class=\"foo\">Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>", "<p class=\"vc foo\">jXzhq RmBrq eXoXz BRc yqhc, PXvBhPchcrz <b>yeRmRBRPRvL</b> hoRc. fhmhooyc, ehohvRcR!</p>"},
		},
		"keep existing class properties by adding new css class",
	},
	{
		"<!-- wp:paragraph -->\n<p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Molestiae minus at aut illo esse unde nemo sint reprehenderit et veritatis qui vel aspernatur sunt, explicabo consequuntur obcaecati similique suscipit dicta! <a rel=\"noreferrer noopener\" href=\"https://example.com/foo/bar\" target=\"_blank\">Link</a> Lorem, ipsum dolor sit amet consectetur adipisicing elit. Laborum quisquam nobis doloremque, est ut veritatis.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Debitis autem aliquam temporibus saepe, nisi suscipit eius accusamus fuga nesciunt porro sequi qui doloremque voluptates doloribus facere maxime vero deleniti similique aut sed nostrum placeat sapiente cumque molestias. Accusantium, possimus sit! Fuga laborum non, quod repellendus inventore iusto commodi nihil, magnam culpa saepe, expedita quibusdam ratione deserunt. Vel voluptas possimus expedita et molestias, dolor ratione! Voluptas facere aperiam labore fugit? Modi nulla ducimus esse rem alias voluptatem eum praesentium consectetur placeat atque omnis architecto, maiores aperiam nihil fugiat magni debitis sint beatae blanditiis quidem harum molestias recusandae! Nostrum asperiores porro dicta nisi debitis quas commodi eaque expedita nam eum animi, quisquam vero dolore officia reiciendis ab magni impedit praesentium voluptatibus deleniti! Vitae quidem consequatur dicta ipsam in ipsum reprehenderit quae accusamus itaque. Architecto aliquam mollitia vel vero, veritatis magnam tempora illo, sint, earum minus explicabo consequuntur similique. Nemo blanditiis expedita nulla?</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Unde, numquam! Voluptatibus, earum, veritatis molestiae assumenda totam nemo accusantium facere labore repellat laudantium deleniti ut distinctio recusandae necessitatibus consequuntur quibusdam! Atque eveniet hic voluptas eos blanditiis dicta explicabo eligendi quis rem a? Voluptate quod in dolorem sequi beatae consequatur laudantium magni.</p>\n<!-- /wp:paragraph -->",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		"vc",
		translations{
			{
				"<p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Molestiae minus at aut illo esse unde nemo sint reprehenderit et veritatis qui vel aspernatur sunt, explicabo consequuntur obcaecati similique suscipit dicta! <a rel=\"noreferrer noopener\" href=\"https://example.com/foo/bar\" target=\"_blank\">Link</a> Lorem, ipsum dolor sit amet consectetur adipisicing elit. Laborum quisquam nobis doloremque, est ut veritatis.</p>",
				"<p class=\"vc\">jXzhq RmBrq eXoXz BRc yqhc PXvBhPchcrz yeRmRBRPRvL hoRc. UXohBcRyh qRvrB yc yrc RooX hBBh rveh vhqX BRvc zhmzhThvehzRc hc lhzRcycRB ArR lho yBmhzvycrz Brvc, hHmoRPyxX PXvBhArrvcrz XxPyhPycR BRqRoRArh BrBPRmRc eRPcy! <a rel=\"noreferrer noopener\" href=\"https://example.com/foo/bar\" target=\"_blank\">jRvC</a> jXzhq, RmBrq eXoXz BRc yqhc PXvBhPchcrz yeRmRBRPRvL hoRc. jyxXzrq ArRBAryq vXxRB eXoXzhqArh, hBc rc lhzRcycRB.</p>",
			},
			{
				"<p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Debitis autem aliquam temporibus saepe, nisi suscipit eius accusamus fuga nesciunt porro sequi qui doloremque voluptates doloribus facere maxime vero deleniti similique aut sed nostrum placeat sapiente cumque molestias. Accusantium, possimus sit! Fuga laborum non, quod repellendus inventore iusto commodi nihil, magnam culpa saepe, expedita quibusdam ratione deserunt. Vel voluptas possimus expedita et molestias, dolor ratione! Voluptas facere aperiam labore fugit? Modi nulla ducimus esse rem alias voluptatem eum praesentium consectetur placeat atque omnis architecto, maiores aperiam nihil fugiat magni debitis sint beatae blanditiis quidem harum molestias recusandae! Nostrum asperiores porro dicta nisi debitis quas commodi eaque expedita nam eum animi, quisquam vero dolore officia reiciendis ab magni impedit praesentium voluptatibus deleniti! Vitae quidem consequatur dicta ipsam in ipsum reprehenderit quae accusamus itaque. Architecto aliquam mollitia vel vero, veritatis magnam tempora illo, sint, earum minus explicabo consequuntur similique. Nemo blanditiis expedita nulla?</p>",
				"<p class=\"vc\">jXzhq RmBrq eXoXz BRc yqhc PXvBhPchcrz, yeRmRBRPRvL hoRc. khxRcRB yrchq yoRAryq chqmXzRxrB Byhmh, vRBR BrBPRmRc hRrB yPPrByqrB irLy vhBPRrvc mXzzX BhArR ArR eXoXzhqArh lXormcychB eXoXzRxrB iyPhzh qyHRqh lhzX ehohvRcR BRqRoRArh yrc Bhe vXBczrq moyPhyc BymRhvch PrqArh qXohBcRyB. aPPrByvcRrq, mXBBRqrB BRc! VrLy oyxXzrq vXv, ArXe zhmhoohverB RvlhvcXzh RrBcX PXqqXeR vRTRo, qyLvyq Promy Byhmh, hHmheRcy ArRxrBeyq zycRXvh ehBhzrvc. Nho lXormcyB mXBBRqrB hHmheRcy hc qXohBcRyB, eXoXz zycRXvh! NXormcyB iyPhzh ymhzRyq oyxXzh irLRc? UXeR vrooy erPRqrB hBBh zhq yoRyB lXormcychq hrq mzyhBhvcRrq PXvBhPchcrz moyPhyc ycArh XqvRB yzPTRchPcX, qyRXzhB ymhzRyq vRTRo irLRyc qyLvR ehxRcRB BRvc xhycyh xoyveRcRRB ArRehq Tyzrq qXohBcRyB zhPrByveyh! QXBczrq yBmhzRXzhB mXzzX eRPcy vRBR ehxRcRB AryB PXqqXeR hyArh hHmheRcy vyq hrq yvRqR, ArRBAryq lhzX eXoXzh XiiRPRy zhRPRhveRB yx qyLvR RqmheRc mzyhBhvcRrq lXormcycRxrB ehohvRcR! NRcyh ArRehq PXvBhArycrz eRPcy RmByq Rv RmBrq zhmzhThvehzRc Aryh yPPrByqrB RcyArh. azPTRchPcX yoRAryq qXooRcRy lho lhzX, lhzRcycRB qyLvyq chqmXzy RooX, BRvc, hyzrq qRvrB hHmoRPyxX PXvBhArrvcrz BRqRoRArh. QhqX xoyveRcRRB hHmheRcy vrooy?</p>",
			},
			{
				"<p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Unde, numquam! Voluptatibus, earum, veritatis molestiae assumenda totam nemo accusantium facere labore repellat laudantium deleniti ut distinctio recusandae necessitatibus consequuntur quibusdam! Atque eveniet hic voluptas eos blanditiis dicta explicabo eligendi quis rem a? Voluptate quod in dolorem sequi beatae consequatur laudantium magni.</p>",
				"<p class=\"vc\">jXzhq RmBrq eXoXz BRc yqhc PXvBhPchcrz, yeRmRBRPRvL hoRc. pveh, vrqAryq! NXormcycRxrB, hyzrq, lhzRcycRB qXohBcRyh yBBrqhvey cXcyq vhqX yPPrByvcRrq iyPhzh oyxXzh zhmhooyc oyreyvcRrq ehohvRcR rc eRBcRvPcRX zhPrByveyh vhPhBBRcycRxrB PXvBhArrvcrz ArRxrBeyq! acArh hlhvRhc TRP lXormcyB hXB xoyveRcRRB eRPcy hHmoRPyxX hoRLhveR ArRB zhq y? NXormcych ArXe Rv eXoXzhq BhArR xhycyh PXvBhArycrz oyreyvcRrq qyLvR.</p>",
			},
		},
		"wordpress blog entry example",
	},
	{
		"",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		"aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO",
		"vc",
		translations{},
		"test empty input",
	},
	// TODO: test brocken xml input
	// TODO: test big input
	// TODO: test html special cases like <br> (invalid xml)
	// TODO: test input with html declaration
	// TODO: test other keys
	// TODO: test reverse translation
}

/**
 * validate functionality of the encoder
 */
func TestEncodeHTML(t *testing.T) {

	for _, testCase := range testCases {

		observed, err := HTML(testCase.input, testCase.keyFrom, testCase.keyTo, testCase.configCSS)
		if err != nil {
			t.Error(err)
		}

		// create expected string from translations
		expected := testCase.input
		for _, translation := range testCase.translations {
			expected = strings.Replace(expected, translation.original, translation.replacement, -1)
		}

		// compare
		if observed != expected {
			t.Errorf("\ninput='%s',\ntestCase.keyFrom='%s',\ntestCase.keyTo='%s',\ntestCase.configCSS='%s',\nobserved='%s',\nexpected='%s',\ndescription: %s", testCase.input, testCase.keyFrom, testCase.keyTo, testCase.configCSS, observed, expected, testCase.description)
		}
	}
}
