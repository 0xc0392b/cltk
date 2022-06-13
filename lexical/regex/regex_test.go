package regex

import "testing"

// tests the regex construction algorithm and whether valid/invalid
// strings are correctly accepted/rejected.
func TestRegex(t *testing.T) {
	for _, test := range []struct {
		pattern     string
		expectedErr bool
		inputs      []regexTestInput
	}{
		{
			pattern:     "a|b",
			expectedErr: false,
			inputs: []regexTestInput{
				{"", false},
				{"c", false},
				{"a", true},
				{"b", true},
			},
		},

		{
			pattern:     "(a|b|c)*",
			expectedErr: false,
			inputs: []regexTestInput{
				{"", true},
				{"abab", true},
				{"ababc", true},
				{"ccccca", true},
				{"ccbccca", true},
				{"ccbcccax", false},
			},
		},

		{
			pattern:     "((a|b|c)|x&y&z)*",
			expectedErr: false,
			inputs: []regexTestInput{
				{"", true},
				{"abab", true},
				{"ababc", true},
				{"ccccca", true},
				{"ccbccca", true},
				{"ccbcccax", false},
				{"ccybcccax", false},
				{"ccybcccaz", false},
				{"xyzxyzaaabbbccc", true},
				{"xyzxyzaaabbbcccx", false},
				{"xyzxyzaaabbbcccy", false},
				{"xyzxyzaaabbbcccxy", false},
				{"xyzxyzaaabbbccczxy", false},
			},
		},

		{
			pattern: "(a|b|c|d|e|f|g|h|i|j|k|l|m|n|o|p|" +
				"q|r|s|t|u|v|w|x|y|z| |.|'|,|:|!|?)*",
			expectedErr: false,
			inputs: []regexTestInput{
				{"hello, regular grammar... ", true},
				{"william santos", true},
				{"another exciting test string", true},
				{"accepts basic punctuation: .,':", true},
				{"Doesn't accept capitals.", false},
				{"remember when adrian said:", true},
				{"'glyphs go in glyphs come out'?", true},
			},
		},
	} {
		if r, err := New(test.pattern); err != nil {
			if !test.expectedErr {
				t.Error(test.pattern, err,
					"but expected no error")
			}
		} else {
			for _, input := range test.inputs {
				switch r.AcceptsStr(input.str) {
				case true:
					if !input.expectedAccept {
						t.Error(test.pattern, "accepted",
							input.str, "but expected reject")
					}

				case false:
					if input.expectedAccept {
						t.Error(test.pattern, "rejected",
							input.str, "but expected accept")
					}
				}
			}
		}
	}
}
