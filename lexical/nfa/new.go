package nfa

// returns a new Alphabet with an initial sequence of symbols.
// an Alphabet is a set of unique Symbols.
func newAlphabet(symbols ...Symbol) Alphabet {
	alpha := make(Alphabet)

	for _, symbol := range symbols {
		alpha[symbol] = true
	}

	return alpha
}

// constructs a new Automaton and returns a pointer to it. sets the
// Automaton's graph to the graph returned by the outer-most
// Thompsoner expression. also sets the Automaton's graph traversal
// algorithm to depth-first search.
func NewAutomaton(exp Thompsoner) *Automaton {
	alpha, graph := exp.Thompson(0)

	a := new(Automaton)
	a.alpha = alpha
	a.graph = graph
	a.acceptsCheck = dfs

	return a
}
