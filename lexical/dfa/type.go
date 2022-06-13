package dfa

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"

// an accepter is a function that takes a slice of NFA sybols and a
// DFA vertex, and tests whether that sequence of symbols forms a
// valid string in the language defined by the DFA graph.
type accepter func(input []nfa.Symbol, start *Vertex) bool

// a set of pointers to NFA vertices.
type nfaVertexSet map[*nfa.Vertex]bool

// an alias for a slice of NFA vertices. satisfies the sortable
// interface defined below.
type nfaVerticesAscendingId []*nfa.Vertex

// a stack of NFA vertices.
type nfaVertexStack struct {
	top *nfaVertexStackNode
}

// a node in an NFA vertex stack.
type nfaVertexStackNode struct {
	vertex *nfa.Vertex
	next   *nfaVertexStackNode
}

// a trie data structure that maps ordered sequences of NFA states
// to the DFA state that contains those NFA states.
type vertexTrie struct {
	root *vertexTrieNode
}

// a node in a vertex trie.
type vertexTrieNode struct {
	children map[int]*vertexTrieNode
	vertex   *Vertex
}

// a type is sortable if it implements the following functions.
type sortable interface {
	Gt(i, j int) bool
	Swap(i, j int)
	Len() int
}

// a DFA automaton has a graph and a function that checks whether strings
// are accepted by that graph.
type Automaton struct {
	graph        *Graph
	acceptsCheck accepter
}

// a DFA vertex has a unique Id and a map of NFA symbols to other DFA
// vertices. the vertex is accepting if its set of NFA vertices contains
// the NFA's accepting state. does not maintain a pointer to the set
// of NFA states computed by the subset construction algorithm because
// they are not required by the acceptsCheck accepter algorithm.
type Vertex struct {
	Id        int
	Accepting bool
	Outgoing  map[nfa.Symbol]*Vertex
}

// a graph simply has a pointer to a start vertex. no need to maintain
// a pointer to the vertex at the "end" of the graph, wherever that
// may be.
type Graph struct {
	Start *Vertex
}
