package opp

// returns true if the top of the stack is a nil pointer.
func (s astNodeStack) Empty() bool {
	return s.top == nil
}

// replaces the top of the stack with a new node and sets
// the new node's next pointer to the previous stack top.
func (s *astNodeStack) Push(node *ASTNode) {
	s.top = &astNodeStackNode{node, s.top}
}

// pops and returns the top of the stack. returns nil if the
// stack is empty. otherwise, replaces the stack top with the
// current top's next node.
func (s *astNodeStack) Pop() *ASTNode {
	if s.top == nil {
		return nil
	} else {
		popped := s.top
		s.top = popped.next
		popped.next = nil
		return popped.node
	}
}

// returns true if the top of the stack is a nil pointer.
func (s tokenStack) Empty() bool {
	return s.top == nil
}

// replaces the top of the stack with a new node and sets
// the new node's next pointer to the previous stack top.
func (s *tokenStack) Push(token Token) {
	s.top = &tokenStackNode{token, s.top}
}

// returns the current stack top without popping it. returns nil
// if the stack is empty.
func (s tokenStack) Peek() Token {
	if s.top == nil {
		return nil
	} else {
		return s.top.token
	}
}

// pops and returns the top of the stack. returns nil if the
// stack is empty. otherwise, replaces the stack top with the
// current top's next node. clears the popped node's next
// pointer.
func (s *tokenStack) Pop() Token {
	if s.top == nil {
		return nil
	} else {
		popped := s.top
		s.top = popped.next
		popped.next = nil
		return popped.token
	}
}
