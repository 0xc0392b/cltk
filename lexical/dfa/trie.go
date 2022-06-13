package dfa

// checks if an NFA vertex set forms a key that exists in a trie.
// first sort the NFA vertices by their Ids and then walk the trie
// trying to reach the end of a chain of NFA vertices. if cannot
// get to the end then the trie does not contain the set.
func (trie vertexTrie) Contains(set nfaVertexSet) bool {
	currentNode := trie.root
	nfaVertices := set.Vertices()

	insertionSort(nfaVerticesAscendingId(nfaVertices))

	for _, vertex := range nfaVertices {
		if currentNode == nil {
			return false
		} else {
			if nextNode, ok := currentNode.
				children[vertex.Id]; ok {

				currentNode = nextNode
			} else {
				return false
			}
		}
	}

	// reached the end so the trie must contain the set
	return true
}

// adds a DFA vertex to the trie using a set of NFA vertices at the
// key. the vertices are sorted by their Id and then added to the
// trie. the last trie node holds the pointer to the given DFA vertex.
func (trie *vertexTrie) Put(set nfaVertexSet, vertex *Vertex) {
	if trie.root == nil {
		trie.root = newVertexTrieNode()
	}

	currentNode := trie.root
	nfaVertices := set.Vertices()

	insertionSort(nfaVerticesAscendingId(nfaVertices))

	for _, vertex := range nfaVertices {
		if child, ok := currentNode.
			children[vertex.Id]; !ok {

			// the next node does not yet exist
			// so add it first.
			next := newVertexTrieNode()
			currentNode.children[vertex.Id] = next
			currentNode = next
		} else {
			// the next node already exists
			// (key is sharing some Ids with other
			// keys).
			currentNode = child
		}
	}

	// store the vertex in the trie
	currentNode.vertex = vertex
}

// gets and returns a DFA vertex set from a trie using an NFA
// vertex set as the key. returns nil if the trie does not
// contain the key.
func (trie vertexTrie) Get(set nfaVertexSet) *Vertex {
	currentNode := trie.root
	nfaVertices := set.Vertices()

	insertionSort(nfaVerticesAscendingId(nfaVertices))

	for _, vertex := range nfaVertices {
		if currentNode == nil {
			return nil
		} else {
			if nextNode, ok := currentNode.
				children[vertex.Id]; ok {

				currentNode = nextNode
			} else {
				// can't go any further -
				// the trie does not contain
				// the key.
				return nil
			}
		}
	}

	// pointer to the vertex associated with the set
	return currentNode.vertex
}
