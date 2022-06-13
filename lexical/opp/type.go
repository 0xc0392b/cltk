package opp

import "errors"

// the error returned when an expression has
// mismatched parentheses.
var MismatchedBrackets = errors.New("expression has mismatched parentheses")

// a stack of AST nodes.
type astNodeStack struct {
	top *astNodeStackNode
}

// a node in an astNodeStack. contains a pointer to an ASTNode and a
// pointer to the next astNodeStackNode if there is one.
type astNodeStackNode struct {
	node *ASTNode
	next *astNodeStackNode
}

// a stack of tokens.
type tokenStack struct {
	top *tokenStackNode
}

// a node in a token stack. has a token and a pointer to the next
// tokenStackNode if there is one.
type tokenStackNode struct {
	token Token
	next  *tokenStackNode
}

// a FIFO queue of tokens. has a pointers to both the top and end of
// the queue for fast insertions and deletions.
type tokenQueue struct {
	top, tail *tokenQueueNode
}

// a node in a tokenQueue. has a token and a pointer to the next
// tokenQueueNode if there is one.
type tokenQueueNode struct {
	token Token
	next  *tokenQueueNode
}

// a helper type for unit tests. simply an alias for a slice of tokens.
type testExpression []Token

// example operand token for unit tests.
type operandToken []rune

// example left bracket token for unit tests.
type leftBracketToken []rune

// example right bracket token for unit tests.
type rightBracketToken []rune

// example operator token for unit tests.
type operatorToken struct {
	lexeme     []rune
	precedence int
}

// a token is a type that implements the function signatures below.
// a token's lexeme may be more than one rune.
type Token interface {
	// TODO: maybe consolidate these?
	Operator() bool
	Function() bool
	LeftBracket() bool
	RightBracket() bool
	Unary() bool
	Binary() bool

	// these are fine
	LeftAssociative() bool
	Precedence() int
	Lexeme() []rune
}

// an abstract syntax tree has a pointer to its root node.
type AST struct {
	Root *ASTNode
}

// a node in an AST has a single token and a list of zero or
// more child nodes.
type ASTNode struct {
	Token    Token
	Children []*ASTNode
}
