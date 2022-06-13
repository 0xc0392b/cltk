package dfa

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"

// returns true if the stack's top is a nil pointer.
func (s nfaVertexStack) Empty() bool {
	return s.top == nil
}

// adds a new stack node to the top of the stack and sets its next node
// pointer to the node previously at the top.
func (s *nfaVertexStack) Push(vertex *nfa.Vertex) {
	s.top = &nfaVertexStackNode{vertex, s.top}
}

// returns nil if the stack is empty. otherwise, replaces the top of
// the stack with the top's next node. then clears the popped node's
// next pointer before returning it.
func (s *nfaVertexStack) Pop() *nfa.Vertex {
	if s.top == nil {
		return nil
	} else {
		popped := s.top
		s.top = popped.next
		popped.next = nil
		return popped.vertex
	}
}
