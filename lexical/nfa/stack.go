package nfa

// a vertexStack is empty if the pointer to the node at the top of
// the stack is a nil pointer.
func (s vertexStack) Empty() bool {
	return s.top == nil
}

// inserts a new node at the top of the stack and sets its next node
// pointer to the node previously at the top (which will always be nil
// if the stack was empty). a node contains a symbol index and a
// pointer to a vertex.
func (s *vertexStack) Push(symbolIndex int, vertex *Vertex) {
	s.top = &vertexStackNode{symbolIndex, vertex, s.top}
}

// pops and returns the symbol index and vertex that was at the top of
// the stack. returns zero and nil if the stack was empty. a pop works
// by temporarily storing the pointer to the top of the stack, and then
// replacing the top with the old top's next node.
func (s *vertexStack) Pop() (int, *Vertex) {
	if s.top == nil {
		return 0, nil
	} else {
		popped := s.top
		s.top = popped.next
		popped.next = nil
		return popped.symbolIndex, popped.vertex
	}
}
