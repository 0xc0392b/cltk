package dfa

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"

// walks a DFA graph given a slice of input symbols. returns
// true if all symbols are consumed and the walk terminates on
// an accepting state. else returns false. unlike the NFA walk
// this algorithm does not use back tracking.
func walkDFA(input []nfa.Symbol, start *Vertex) bool {
	currentVertex := start
	symbolIndex := 0
	stuck := false

	// keep walking until run out of symbols or stuck at
	// a vertex (cannot transition on current symbol).
	for symbolIndex < len(input) && !stuck {
		symbol := input[symbolIndex]

		if nextVertex, ok :=
			currentVertex.Outgoing[symbol]; ok {

			currentVertex = nextVertex
			symbolIndex += 1
		} else {
			stuck = true
		}
	}

	// the string was accepted if the automaton reached
	// an accepting state without getting stuck.
	if currentVertex.Accepting && !stuck {
		return true
	} else {
		return false
	}
}
