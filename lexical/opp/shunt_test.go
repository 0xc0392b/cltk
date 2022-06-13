package opp

import "testing"

// tests the shunting yard algorithm with a series of test expressions.
// these tests work by comparing infix expressions to their equivalent
// postfix representations.
func TestShuntingYard(t *testing.T) {
	for _, test := range []struct {
		expression     testExpression
		expectedErr    bool
		expectedOutput string
	}{
		// infix   : 1 + 2 + 3
		// postfix : 1 2 3 + +
		{
			expression:     strToExpression("1+2+3"),
			expectedErr:    false,
			expectedOutput: "12+3+",
		},

		// infix   : 9 - 1 + 2 + 3 * 4 / 2
		// postfix : 9 1 - 2 + 3 4 * 2 / +
		{
			expression:     strToExpression("9-1+2+3*4/2"),
			expectedErr:    false,
			expectedOutput: "91-2+34*2/+",
		},

		// infix   : 1 + 2 + 3 * (4 / 2) - 1 * 6 / 3
		// postfix : 1 2 + 3 4 2 / * + 1 6 * 3 / -
		{
			expression:     strToExpression("1+2+3*(4/2)-1*6/3"),
			expectedErr:    false,
			expectedOutput: "12+342/*+16*3/-",
		},

		// infix   : 1 + 2 * (
		// postfix : mismatched parentheses
		{
			expression:  strToExpression("1+2*("),
			expectedErr: true,
		},

		// infix   : ) + 1 * 2
		// postfix : mismatched parentheses
		{
			expression:  strToExpression(")+1*2"),
			expectedErr: true,
		},

		// infix   : 1 + 2 + 3 + 4 * (1 - (8 / 4) / 2
		// postfix : mismatched parentheses
		{
			expression:  strToExpression("1+2+3+4*(1-(8/4)/2"),
			expectedErr: true,
		},
	} {
		if output, err :=
			shuntingYard(test.expression); err != nil {

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

			for !output.Empty() {
				popped := output.Pop()
				characters += string(popped.Lexeme())
			}

			if characters != test.expectedOutput {
				t.Error("got", characters,
					"expected", test.expectedOutput)
			}
		}
	}
}
