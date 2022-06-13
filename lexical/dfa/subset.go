package dfa

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"

// the subset construction algorithm. takes an NFA alphabet, NFA accepting
// vertex, and set of NFA vertices. uses a trie to map sequences of ordered
// NFA state Ids to their respective DFA state set. recursively computes
// subset construction until the symbol colsure of a given epsilon closure
// returns the empty set.
func subset(alpha nfa.Alphabet, counter int,
	target *nfa.Vertex, trie *vertexTrie,
	set nfaVertexSet) *Graph {

	// the epsilon closure is the current set union-ed with the set of
	// vertices reachable from the current set by only following
	// epsilon transitions.
	epsilonClosure := set.Union(closure(set, nfa.Epsilon))

	if !trie.Contains(epsilonClosure) {
		// if the above epsilon closure is unique (the NFA vertices
		// ordered by their Ids does not already exist in the trie),
		// then create and add a new DFA state to the trie. that state
		// is an accepting state if it contains the NFA's accepting
		// state. recursively find the other DFA states by computing
		// the closure for each symbol in the alphabet on the current
		// epsilon closure.
		vertex := &Vertex{counter,
			epsilonClosure.Contains(target),
			make(map[nfa.Symbol]*Vertex)}

		trie.Put(epsilonClosure, vertex)

		for _, symbol := range alpha.Symbols() {
			if symbolClosure := closure(epsilonClosure,
				symbol); !symbolClosure.Empty() {

				// compute the subgraph and "attach" it to the
				// current vertex with an edge laballed with
				// the current symbol.
				subgraph := subset(alpha,
					counter+len(vertex.Outgoing),
					target, trie, symbolClosure)

				vertex.Outgoing[symbol] = subgraph.Start
			}
		}

		return &Graph{vertex}
	} else {
		// if the above epsilon closure is not unique, then get and
		// return its corresponding DFA set from the trie.
		return &Graph{trie.Get(epsilonClosure)}
	}
}
