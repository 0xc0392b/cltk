package opp

func (t operandToken) Lexeme() []rune {
	return []rune(t)
}

func (t operandToken) Operator() bool {
	return false
}

func (t operandToken) Precedence() int {
	return 0
}

func (t operandToken) LeftAssociative() bool {
	return false
}

func (t operandToken) Function() bool {
	return false
}

func (t operandToken) LeftBracket() bool {
	return false
}

func (t operandToken) RightBracket() bool {
	return false
}

func (t operandToken) Unary() bool {
	return false
}

func (t operandToken) Binary() bool {
	return false
}

func (t operatorToken) Lexeme() []rune {
	return t.lexeme
}

func (t operatorToken) Operator() bool {
	return true
}

func (t operatorToken) Precedence() int {
	return t.precedence
}

func (t operatorToken) LeftAssociative() bool {
	return true
}

func (t operatorToken) Function() bool {
	return false
}

func (t operatorToken) LeftBracket() bool {
	return false
}

func (t operatorToken) RightBracket() bool {
	return false
}

func (t operatorToken) Unary() bool {
	return false
}

func (t operatorToken) Binary() bool {
	return true // change this
}

func (t leftBracketToken) Lexeme() []rune {
	return []rune(t)
}

func (t leftBracketToken) Operator() bool {
	return false
}

func (t leftBracketToken) Precedence() int {
	return 0
}

func (t leftBracketToken) LeftAssociative() bool {
	return false
}

func (t leftBracketToken) Function() bool {
	return false
}

func (t leftBracketToken) LeftBracket() bool {
	return true
}

func (t leftBracketToken) RightBracket() bool {
	return false
}

func (t leftBracketToken) Unary() bool {
	return false
}

func (t leftBracketToken) Binary() bool {
	return false
}

func (t rightBracketToken) Lexeme() []rune {
	return []rune(t)
}

func (t rightBracketToken) Operator() bool {
	return false
}

func (t rightBracketToken) Precedence() int {
	return 0
}

func (t rightBracketToken) LeftAssociative() bool {
	return false
}

func (t rightBracketToken) Function() bool {
	return false
}

func (t rightBracketToken) LeftBracket() bool {
	return false
}

func (t rightBracketToken) RightBracket() bool {
	return true
}

func (t rightBracketToken) Unary() bool {
	return false
}

func (t rightBracketToken) Binary() bool {
	return false
}
