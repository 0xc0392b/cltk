package dfa

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"

// construct and return a new vertex set from zero or more initial
// vertices.
func newVertexSet(vertices ...*nfa.Vertex) nfaVertexSet {
	set := make(nfaVertexSet)

	for _, vertex := range vertices {
		set[vertex] = true
	}

	return set
}

// construct and return a new vertex stack with zero or more vertices
// already pushed onto it.
func newNFAVertexStack(vertices ...*nfa.Vertex) *nfaVertexStack {
	stack := new(nfaVertexStack)

	for _, vertex := range vertices {
		stack.Push(vertex)
	}

	return stack
}

// construct and return a new vertex trie node with an initial
// empty map of child nodes.
func newVertexTrieNode() *vertexTrieNode {
	node := new(vertexTrieNode)
	node.children = make(map[int]*vertexTrieNode)

	return node
}

// construct and return a new DFA automaton. the DFA's graph is
// represented using a graph of states, where each state is a set of
// NFA states. the DFA is built from an NFA using the subset
// construction algorithm.
func NewAutomaton(nfa *nfa.Automaton) *Automaton {
	nfaGraph := nfa.Graph()

	a := new(Automaton)
	a.graph = subset(nfa.Alphabet(), 0, nfaGraph.End,
		&vertexTrie{}, newVertexSet(nfaGraph.Start))
	a.acceptsCheck = walkDFA

	return a
}
