package nfa

// symbol for epsilon transition
const Epsilon = Symbol(0)

// a accepter is a function. it tests whether a sequence of symbols form
// a string defined by a regular grammar that is encoded as a graph.
type accepter func(input []Symbol, start, target *Vertex) bool

// a linked list of vertexStackNode structs with the following stack
// functions:
//   1. Empty()
//   2. Push()
//   3. Pop()
type vertexStack struct {
	top *vertexStackNode
}

// a node in a vertex stack. contains the index of a character in an input
// string, a vertex in a graph, and a pointer to the next node in the list
// if there is one.
type vertexStackNode struct {
	symbolIndex int
	vertex      *Vertex
	next        *vertexStackNode
}

// represents a symbol in an input string. also satisfies the Thompsoner
// interface: can return a Thompson's construction accept-symbol
// graph for the rune it wraps.
type Symbol rune

// Alpabet is a unique set of Symbols
type Alphabet map[Symbol]bool

// represents a Kleene closure. this is a unary operator so it only has one
// sub-expression. satisfies the Thompsoner interface: can return a
// Thompson's construction Kleene star graph for the sub-expression, which
// also satisfies Thompsoner.
type Closure struct {
	Exp Thompsoner
}

// represents a union. this is a binary operator so it has two
// sub-expressions. satisfies the Thompsoner interface: can return a
// Thompson's construction union graph for the two sub-expressions, which
// also satisfy Thompsoner.
type Union struct {
	LeftExp, RightExp Thompsoner
}

// represents a concatenation. this is a binary operator so it has two
// sub-expressions. satisfies the Thompsoner interface: can return a
// Thompson's construction concatenation graph for the two sub-expressions,
// which also satisfy Thompsoner.
type Concat struct {
	LeftExp, RightExp Thompsoner
}

// a Thompsoner is anything that can return a Grah representation of itself
// using Thompson's construction.
type Thompsoner interface {
	Thompson(counter int) (Alphabet, *Graph)
}

// an Automaton defines a regular language using a regular grammar encoded as
// a graph. to walk the graph, to test if a string belongs to the language,
// Automaton uses an accepter function which implements a suitable graph
// traversal algorithm.
type Automaton struct {
	alpha        Alphabet
	graph        *Graph
	acceptsCheck accepter
}

// a Vertex has a unique id and a list of out-going edges.
type Vertex struct {
	Id       int
	Outgoing []Edge
}

// an edge has a pointer to a destination vertex and a label: which is the
// symbol that the edge "accepts". if the symbol can either be a symbol in
// the language's alphabet or epsilon (null transition).
type Edge struct {
	Dest    *Vertex
	Accepts Symbol
}

// a Graph has pointers to start and end vertices.
type Graph struct {
	Start, End *Vertex
}
