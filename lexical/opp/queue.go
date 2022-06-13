package opp

// returns true if a token queue has no top.
func (s tokenQueue) Empty() bool {
	return s.top == nil
}

// adds a new token to a token queue. if the queue is empty then the
// start and end pointers are set to the new node. otherwise, the new
// node is set to the current tail's "next" node.
func (s *tokenQueue) Push(token Token) {
	node := new(tokenQueueNode)
	node.token = token

	if s.top == nil {
		s.top = node
		s.tail = node
	} else {
		s.tail.next = node
		s.tail = node
	}
}

// pops a token from the front of the queue. returns nil if the queue
// is empty. otherwise, sets the top of the queue to the current top's
// "next" node before returning the popped token. the popped node's next
// pointer it also cleared.
func (s *tokenQueue) Pop() Token {
	if s.top == nil {
		return nil
	} else {
		popped := s.top
		s.top = popped.next
		popped.next = nil
		return popped.token
	}
}
