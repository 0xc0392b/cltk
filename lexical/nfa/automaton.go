package nfa

// returns a pointer to the Automaton's underlying graph.
func (a Automaton) Graph() *Graph {
	return a.graph
}

// returns the Automaton's alphabet.
func (a Automaton) Alphabet() Alphabet {
	return a.alpha
}

// returns true if the given string belongs to the language defined
// by the Automaton, else returns false.
func (a Automaton) AcceptsStr(input string) bool {
	if a.acceptsCheck([]Symbol(input),
		a.graph.Start, a.graph.End) {
		return true
	} else {
		return false
	}
}
