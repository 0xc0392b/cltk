package regex

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/opp"

// translates a string of runes into a slice of tokens. any rune that
// isn't an operator is assumed to be an operand.
func strToExpression(input string) []opp.Token {
	expression := make([]opp.Token, len(input))

	for i := 0; i < len(input); i++ {
		switch char := rune(input[i]); char {
		// union operator
		case '|':
			expression[i] =
				unionToken([]rune{char})

		// concatenate operator
		case '&':
			expression[i] =
				concatToken([]rune{char})

		// closure operator
		case '*':
			expression[i] =
				closureToken([]rune{char})

		// opening brackets
		case '(':
			expression[i] =
				leftBracketToken([]rune{char})

		// closing brackets
		case ')':
			expression[i] =
				rightBracketToken([]rune{char})

		// by default assume operand
		default:
			expression[i] =
				operandToken([]rune{char})
		}
	}

	return expression
}
