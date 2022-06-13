package tokeniser

// returns the string representation of a token.
func (t Token) String() string {
	return "[" + t.Name + ":" + string(t.Lexeme) + "]"
}
