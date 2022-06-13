package opp

// implements Dijkstra's shunting yard algorithm. in this context,
// it is used as an operator precedence parser.
func shuntingYard(tokens []Token) (*tokenQueue, error) {
	// postfix expression is built here
	output := tokenQueue{}

	// the operator stack
	operators := tokenStack{}

	for _, token := range tokens {
		switch {
		// the token is an operator
		case token.Operator():
			// whilst the operator at the top of the operator
			// stack has a higher precedence than the current
			// token, pop the operator stack into the output
			// queue.
			for !operators.Empty() {
				top := operators.Peek()

				if top.Precedence() >
					token.Precedence() {

					output.Push(top)
					operators.Pop()

				} else if top.Precedence() ==
					token.Precedence() &&
					token.LeftAssociative() {

					output.Push(top)
					operators.Pop()

				} else {
					break
				}
			}

			// always push the current token onto the operator
			// stack.
			operators.Push(token)

		// the token is a right bracket
		case token.RightBracket():
			// assume the expression has mismatched
			// brackets.
			mismatched := true

			// while the expression is mismatched, pop from the
			// operator stack and append to the output queue.
			// stop when left bracket is read.
			for !operators.Empty() && mismatched {
				popped := operators.Pop()

				if popped.LeftBracket() {
					mismatched = false
				} else {
					output.Push(popped)
				}
			}

			switch {
			// the expression has mismatched brackets so return
			// an error.
			case mismatched:
				return nil, MismatchedBrackets

			// if the top of the stack is a function, pop it from
			// the stack into the output queue.
			case !operators.Empty():
				top := operators.Peek()

				if top.Function() {
					output.Push(top)
					operators.Pop()
				}
			}

		// the token is a left bracket
		case token.LeftBracket():
			operators.Push(token)

		// the token is a function
		case token.Function():
			operators.Push(token)

		// by default assume token is an operand
		default:
			output.Push(token)
		}
	}

	// pop all remaining operators from operator stack
	for !operators.Empty() {
		// but if there is still a left bracket on the stack then
		// return a mismatched brackets error.
		if operator := operators.Pop(); operator.LeftBracket() {
			return nil, MismatchedBrackets
		} else {
			output.Push(operator)
		}
	}

	return &output, nil
}
