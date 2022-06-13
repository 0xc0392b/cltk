package regex

// token for concat operator lexeme.
type concatToken []rune

// token for union operator lexeme.
type unionToken []rune

// token for closure operator lexeme.
type closureToken []rune

// token for left bracket lexeme.
type leftBracketToken []rune

// token for right bracket lexeme.
type rightBracketToken []rune

// token for operand lexeme.
type operandToken []rune

// helper type for unit tests.
type regexTestInput struct {
	str            string
	expectedAccept bool
}

// any regularGrammar needs to be able to test whether the
// language described by the grammar contains a particular
// string.
type regularGrammar interface {
	AcceptsStr(input string) bool
}

// a regex has a grammar which satisfies the regularGrammar
// interface.
type Regex struct {
	grammar regularGrammar
}
