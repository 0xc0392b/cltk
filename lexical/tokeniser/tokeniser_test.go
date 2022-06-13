package tokeniser

import (
	"io/ioutil"
	"os"
	"testing"
)

// tests the tokeniser using a set of tokens and multiple input strings.
// simply checks whether the correct number of tokens is returned.
func TestTokeniser(t *testing.T) {
	tokeniser, err := New(
		TokenDef{"WORD",
			"(a|b|c|d|e|f|g|h|i|j|k|l|m|n|o|p|q|r|s|t|u|v|w|x|y|z|" +
				"A|B|C|D|E|F|G|H|I|J|K|L|M|N|O|P|Q|R|S|T|U|V|W|" +
				"X|Y|W|Z|0|1|2|3|4|5|6|7|8|9|-|')*"},

		TokenDef{"CITATION",
			"[&(0|1|2|3|4|5|6|7|8|9)*&]"},

		TokenDef{"PUNCTUATION",
			".|'|,|:|;|!|?"},

		TokenDef{"WHITE_SPACE",
			" ( )*"},
	)
	if err != nil {
		t.Error(err, "but expected no error")
	}

	for _, test := range []struct {
		name               string
		path               string
		expectedErr        bool
		expectedTokenCount int
	}{
		{
			name:               "climate change",
			path:               "./test_corpuses/climatechange.txt",
			expectedErr:        false,
			expectedTokenCount: 543,
		},

		{
			name:               "lorem ipsum",
			path:               "./test_corpuses/loremipsum.txt",
			expectedErr:        false,
			expectedTokenCount: 38811,
		},
	} {
		if file, err := os.Open(test.path); err != nil {
			t.Error(test.name,
				"could not open corpus file", err)
		} else {
			defer func() {
				if err := file.Close(); err != nil {
					t.Error(test.name,
						"could not close corpus file", err)
				}
			}()

			if bytes, err := ioutil.ReadAll(file); err != nil {
				t.Error(test.name,
					"could not read from corpus file", err)
			} else {
				if tokens, err := tokeniser.TokeniseStr(
					string(bytes)); err != nil {

					if !test.expectedErr {
						t.Error(test.name,
							err, "but expected no error")
					}
				} else {
					if count := len(tokens); count !=
						test.expectedTokenCount {

						t.Error(test.name, count, "but expected",
							test.expectedTokenCount)
					}
				}
			}
		}
	}
}
