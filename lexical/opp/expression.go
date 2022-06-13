package opp

// converts a sequence of runes in a string into a testExpression.
// this is a helper function for unit tests. only accepts single
// digit numbers e.g. 1+2+3*(4/2) and not (200-100)/2.
func strToExpression(input string) testExpression {
	// the expression will contain the number of runes in the
	// input string.
	expression := make(testExpression, len(input))

	for i := 0; i < len(input); i++ {
		// iterate over all runes in the input and
		// extract operators.
		switch char := rune(input[i]); char {
		// addition and subtraction have the same precedence
		case '+':
			fallthrough
		case '-':
			expression[i] =
				operatorToken{[]rune{char}, 1}

		// multiply and divide have the same precedence
		case '*':
			fallthrough
		case '/':
			expression[i] =
				operatorToken{[]rune{char}, 2}

		// opening (left) parenthesis
		case '(':
			expression[i] =
				leftBracketToken([]rune{char})

		// closing (right) parenthesis
		case ')':
			expression[i] =
				rightBracketToken([]rune{char})

		// if not an operator then the rune must be an
		// operand.
		default:
			expression[i] =
				operandToken([]rune{char})
		}
	}

	return expression
}
