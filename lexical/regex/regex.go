package regex

// tests whether a string is accepted by a regular expression
// using the underlying regular grammar's AcceptsStr function.
func (r Regex) AcceptsStr(input string) bool {
	return r.grammar.AcceptsStr(input)
}
