package opp

import "testing"

// tests the parse algorithm with a series of test expressions -- testing
// both the underlting parsing algorithm and the AST construction algorithm.
// these tests work by comparing infix expressions to their equivalent
// prefix representations. the parsing algorithm turns the infix
// expression into a postfix one. the AST construction algorithm turns
// the postfix expression into a prefix one.
func TestParseTokens(t *testing.T) {
	for _, test := range []struct {
		expression     testExpression
		expectedErr    bool
		expectedOutput string
	}{
		// infix   : 1 + 2 + 3
		// postfix : 1 2 3 + +
		// prefix  : + + 1 2 3
		{
			expression:     strToExpression("1+2+3"),
			expectedErr:    false,
			expectedOutput: "++123",
		},

		// infix   : 9 - 1 + 2 + 3 * 4 / 2
		// postfix : 9 1 - 2 + 3 4 * 2 / +
		// prefix  : + + - 9 1 2 / * 3 4 2
		{
			expression:     strToExpression("9-1+2+3*4/2"),
			expectedErr:    false,
			expectedOutput: "++-912/*342",
		},

		// infix   : 1 + 2 + 3 * (4 / 2) - 1 * 6 / 3
		// postfix : 1 2 + 3 4 2 / * + 1 6 * 3 / -
		// prefix  : - + + 1 2 * 3 / 4 2 / * 1 6 3
		{
			expression:     strToExpression("1+2+3*(4/2)-1*6/3"),
			expectedErr:    false,
			expectedOutput: "-++12*3/42/*163",
		},

		// infix   : 1 + 2 * (
		// prefix  : mismatched parentheses
		{
			expression:  strToExpression("1+2*("),
			expectedErr: true,
		},

		// infix   : ) + 1 * 2
		// prefix  : mismatched parentheses
		{
			expression:  strToExpression(")+1*2"),
			expectedErr: true,
		},

		// infix   : 1 + 2 + 3 + 4 * (1 - (8 / 4) / 2
		// prefix  : mismatched parentheses
		{
			expression:  strToExpression("1+2+3+4*(1-(8/4)/2"),
			expectedErr: true,
		},
	} {
		if ast, err := ParseTokens(test.expression); err != nil {
			if !test.expectedErr {
				characters := ""

				for _, token := range test.expression {
					characters += string(token.Lexeme())
				}

				t.Error(characters, err,
					"but expected no error")
			}
		} else {
			characters := ""

			ast.Walk(func(token Token) {
				characters += string(token.Lexeme())
			})

			if characters != test.expectedOutput {
				t.Error("got", characters,
					"expected", test.expectedOutput)
			}
		}
	}
}
