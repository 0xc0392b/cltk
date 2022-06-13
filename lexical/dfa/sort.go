package dfa

// compare two vertices' Ids. returns true if i has a higher Id
// than j, otherwise false.
func (vertices nfaVerticesAscendingId) Gt(i, j int) bool {
	return vertices[i].Id > vertices[j].Id
}

// swaps two vertices by index.
func (vertices nfaVerticesAscendingId) Swap(i, j int) {
	vertices[i], vertices[j] = vertices[j], vertices[i]
}

// returns the number of vertices to be sorted.
func (vertices nfaVerticesAscendingId) Len() int {
	return len(vertices)
}

// implements the insertion sort algorithm over all types that
// satisfy the "sortable" interface defined in this package.
func insertionSort(items sortable) {
	i := 1

	for i < items.Len() {
		j := i

		for j > 0 && items.Gt(j-1, j) {
			items.Swap(j, j-1)
			j -= 1
		}

		i += 1
	}
}
