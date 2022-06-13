package nfa

// implements a stack-based depth-first search over a NFA graph.
// returns true if it can reach the target from the given start and
// consumed all of the input symbols in the process. otherwise returns
// false (=the input string is not accepted by the NFA).
func dfs(input []Symbol, start, target *Vertex) bool {
	stack := vertexStack{}
	accepted := false

	stack.Push(0, start)

	for !(stack.Empty() || accepted) {
		symbolIndex, vertex := stack.Pop()

		if vertex == target && symbolIndex == len(input) {
			accepted = true
		} else {
			for _, edge := range vertex.Outgoing {
				if edge.Accepts == Epsilon {
					stack.Push(symbolIndex,
						edge.Dest)
				} else if symbolIndex < len(input) {
					if edge.Accepts == input[symbolIndex] {
						stack.Push(symbolIndex+1,
							edge.Dest)
					}
				}
			}
		}
	}

	return accepted
}
