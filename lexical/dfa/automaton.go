package dfa

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"

// returns the DFA automaton's underlying graph structure.
func (a Automaton) Graph() *Graph {
	return a.graph
}

// uses the automaton's acceptCheck function to test whether a
// string exists in the language described by the DFA's grammar.
func (a Automaton) AcceptsStr(input string) bool {
	if a.acceptsCheck([]nfa.Symbol(input), a.graph.Start) {
		return true
	} else {
		return false
	}
}

// creates and returns the jump table representation of the DFA
// so it can be saved or inlined into a parser.
func (a Automaton) JumpTable() [][]int {

	// TODO

	return nil
}
