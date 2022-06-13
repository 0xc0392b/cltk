package regex

func (t concatToken) Operator() bool {
	return true
}

func (t concatToken) Function() bool {
	return false
}

func (t concatToken) LeftBracket() bool {
	return false
}

func (t concatToken) RightBracket() bool {
	return false
}

func (t concatToken) Unary() bool {
	return false
}

func (t concatToken) Binary() bool {
	return true
}

func (t concatToken) LeftAssociative() bool {
	return true
}

func (t concatToken) Precedence() int {
	return 2
}

func (t concatToken) Lexeme() []rune {
	return t
}

func (t unionToken) Operator() bool {
	return true
}

func (t unionToken) Function() bool {
	return false
}

func (t unionToken) LeftBracket() bool {
	return false
}

func (t unionToken) RightBracket() bool {
	return false
}

func (t unionToken) Unary() bool {
	return false
}

func (t unionToken) Binary() bool {
	return true
}

func (t unionToken) LeftAssociative() bool {
	return true
}

func (t unionToken) Precedence() int {
	return 1
}

func (t unionToken) Lexeme() []rune {
	return t
}

func (t closureToken) Operator() bool {
	return true
}

func (t closureToken) Function() bool {
	return false
}

func (t closureToken) LeftBracket() bool {
	return false
}

func (t closureToken) RightBracket() bool {
	return false
}

func (t closureToken) Unary() bool {
	return true
}

func (t closureToken) Binary() bool {
	return false
}

func (t closureToken) LeftAssociative() bool {
	return false
}

func (t closureToken) Precedence() int {
	return 3
}

func (t closureToken) Lexeme() []rune {
	return t
}

func (t leftBracketToken) Operator() bool {
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

func (t leftBracketToken) LeftAssociative() bool {
	return false
}

func (t leftBracketToken) Precedence() int {
	return 0
}

func (t leftBracketToken) Lexeme() []rune {
	return t
}

func (t rightBracketToken) Operator() bool {
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

func (t rightBracketToken) LeftAssociative() bool {
	return false
}

func (t rightBracketToken) Precedence() int {
	return 0
}

func (t rightBracketToken) Lexeme() []rune {
	return t
}

func (t operandToken) Operator() bool {
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

func (t operandToken) LeftAssociative() bool {
	return false
}

func (t operandToken) Precedence() int {
	return 0
}

func (t operandToken) Lexeme() []rune {
	return t
}
