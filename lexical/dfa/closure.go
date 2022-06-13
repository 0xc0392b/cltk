package dfa

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"

// finds the closure of a set given a particular input symbol. that
// is, the set of nodes reachable from each node in the set if only
// edges labelled with the input symbol are walked.
func closure(set nfaVertexSet, symbol nfa.Symbol) nfaVertexSet {
	vertices := set.Vertices()
	closureSet := newVertexSet()
	stack := newNFAVertexStack(vertices...)

	// perform a stack-based depth-first search
	for !stack.Empty() {
		vertex := stack.Pop()

		for _, edge := range vertex.Outgoing {
			if edge.Accepts == symbol &&
				!closureSet.Contains(edge.Dest) {

				closureSet.Add(edge.Dest)
				stack.Push(edge.Dest)
			}
		}
	}

	return closureSet
}
