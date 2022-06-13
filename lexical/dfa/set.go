package dfa

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"

// add a new vertex to the vertex set. nothing happens if the vertex
// is already in the set.
func (set nfaVertexSet) Add(vertex *nfa.Vertex) {
	set[vertex] = true
}

// returns true if the set is empty otherwise false.
func (set nfaVertexSet) Empty() bool {
	if len(set) == 0 {
		return true
	} else {
		return false
	}
}

// returns true if the set contains a given vertex otherwise false.
func (set nfaVertexSet) Contains(vertex *nfa.Vertex) bool {
	if _, ok := set[vertex]; ok {
		return true
	} else {
		return false
	}
}

// returns all vertices in the set as a slice. no particular ordering
// should be expected.
func (set nfaVertexSet) Vertices() []*nfa.Vertex {
	vertices := make([]*nfa.Vertex, 0)

	for vertex := range set {
		vertices = append(vertices, vertex)
	}

	return vertices
}

// creates a new set, then adds all vertices in the current set and in
// an "other" set, producing the union of both sets.
func (set nfaVertexSet) Union(other nfaVertexSet) nfaVertexSet {
	newSet := make(nfaVertexSet)

	for vertex := range set {
		newSet[vertex] = true
	}

	for vertex := range other {
		newSet[vertex] = true
	}

	return newSet
}
