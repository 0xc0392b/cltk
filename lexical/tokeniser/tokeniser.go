package tokeniser

// turns an input string into a sequence of tokens. uses the longest
// match strategy to test whether substrings in the input are accepted
// by any of the tokens' string accepting machines.
func (t Tokeniser) TokeniseStr(input string) ([]Token, error) {
	// initialise substring pointers to zero
	i, j := 0, 0

	// slice of tokens is empty to begin with
	tokens := make([]Token, 0)

	// keep track of the longest matching machine (token)
	var longestMatch *machine

	// repeat until j pointer has reached the end of the input
	// string.
	for j < len(input) {
		// assume the current substring has not been accepted
		// by any machine.
		matched := false

		// find which machine first accepts the substring,
		// if any.
		for _, m := range t.machines {
			if m.grammar.AcceptsStr(input[i : j+1]) {
				longestMatch = &m
				matched = true
				break
			}
		}

		if matched {
			// if the substring was accepted then see if a longer
			// version will be accepted.
			j += 1
		} else {
			// the substring was not accepted
			if longestMatch == nil {
				// if it was never accepted in the first place
				// then just try a longer substring until j
				// reaches the end of the input string.
				j += 1
			} else {
				// if the previous (shorter) substring was
				// accepted then save it along with the name of
				// the token it belongs to.
				tokens = append(tokens,
					Token{longestMatch.name,
						[]rune(input[i:j])})

				// then push i forward to the end of the substring
				// and start again.
				i = j
				longestMatch = nil
			}
		}
	}

	return tokens, nil
}
