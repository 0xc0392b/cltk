package opp

// converts a sequence of symbol and operator tokens in postfix
// notation into an abstract syntax tree using a stack-based
// "reverse polish notation evaluation" style algorithm.
func postfixToAST(postfix *tokenQueue) *ASTNode {
	// push and pop to this stack from the postfix token queue
	stack := astNodeStack{}

	for !postfix.Empty() {
		// consume from the queue of tokens until the queue is
		// empty.
		token := postfix.Pop()

		switch {
		// the current token is an operator
		case token.Operator():
			switch {
			// it is an operator with one operand
			case token.Unary():
				operand := stack.Pop()

				node := new(ASTNode)
				node.Token = token
				node.Children = []*ASTNode{
					operand}

				stack.Push(node)

			// it is an operator with two operands
			case token.Binary():
				rOperand := stack.Pop()
				lOperand := stack.Pop()

				node := new(ASTNode)
				node.Token = token
				node.Children = []*ASTNode{
					lOperand, rOperand}

				stack.Push(node)
			}

		// the token is a function, which is the same as a
		// unary operator.
		case token.Function():
			operand := stack.Pop()

			node := new(ASTNode)
			node.Token = token
			node.Children = []*ASTNode{
				operand}

			stack.Push(node)

		// assume the token is just an operand
		default:
			node := new(ASTNode)
			node.Token = token

			stack.Push(node)
		}
	}

	// the last node in the stack will be the root of the AST
	return stack.Pop()
}

// walks the tree with a given function.
func (ast *AST) Walk(do func(token Token)) {
	if ast.Root != nil {
		ast.Root.PreOrderTraversal(do)
	}
}

// do a pre-order traversal of the AST and execute a "do" lambda
// function when visiting each node.
func (node *ASTNode) PreOrderTraversal(do func(token Token)) {
	// call do lambda before visiting child nodes
	do(node.Token)

	// visit each child in order
	for _, child := range node.Children {
		child.PreOrderTraversal(do)
	}
}
