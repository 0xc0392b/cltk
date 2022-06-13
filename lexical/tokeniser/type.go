package tokeniser

// a token's regular grammar tests whether a string belongs
// to the grammar's language.
type tokenRegularGrammar interface {
	AcceptsStr(input string) bool
}

// an accepting machine has a name and a regular grammar.
type machine struct {
	name    string
	grammar tokenRegularGrammar
}

// a token definition has a unique name and a regex pattern
// for matching strings that belong to the token.
type TokenDef struct {
	Name, Pattern string
}

// the name of a token plus its lexeme. a token's lexeme can
// be multiple runes.
type Token struct {
	Name   string
	Lexeme []rune
}

// a tokeniser has multiple accepting machines and uses the
// longest match strategy.
type Tokeniser struct {
	machines []machine
}
