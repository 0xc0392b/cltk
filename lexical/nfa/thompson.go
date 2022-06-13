package nfa

// returns a graph representation of a Kleene closure using
// Thompson's construction. Closure is a Thompsoner and has
// one sub-expression.
func (c Closure) Thompson(counter int) (Alphabet, *Graph) {
	start := new(Vertex)
	end := new(Vertex)
	alpha, exp := c.Exp.Thompson(counter + 1)

	start.Id = counter
	end.Id = exp.End.Id + 1

	start.Outgoing = []Edge{
		Edge{exp.Start, Epsilon},
		Edge{end, Epsilon}}

	exp.End.Outgoing = []Edge{
		Edge{exp.Start, Epsilon},
		Edge{end, Epsilon}}

	return alpha, &Graph{start, end}
}

// returns a graph representation of a union using Thompson's
// construction. Union is a Thompsoner and has two
// sub-expressions.
func (u Union) Thompson(counter int) (Alphabet, *Graph) {
	start := new(Vertex)
	end := new(Vertex)
	lAlpha, lExp := u.LeftExp.Thompson(counter + 1)
	rAlpha, rExp := u.RightExp.Thompson(lExp.End.Id + 1)
	alpha := lAlpha.Union(rAlpha)

	start.Id = counter
	end.Id = rExp.End.Id + 1

	lExp.End.Outgoing = []Edge{
		Edge{end, Epsilon}}

	rExp.End.Outgoing = []Edge{
		Edge{end, Epsilon}}

	start.Outgoing = []Edge{
		Edge{lExp.Start, Epsilon},
		Edge{rExp.Start, Epsilon}}

	return alpha, &Graph{start, end}
}

// returns a graph representation of a concatenation using
// Thompson's construction. Concat is a Thompsoner and has two
// sub-expressions.
func (c Concat) Thompson(counter int) (Alphabet, *Graph) {
	lAlpha, lExp := c.LeftExp.Thompson(counter)
	rAlpha, rExp := c.RightExp.Thompson(lExp.End.Id + 1)
	alpha := lAlpha.Union(rAlpha)

	lExp.End.Outgoing = []Edge{
		Edge{rExp.Start, Epsilon}}

	return alpha, &Graph{lExp.Start, rExp.End}
}

// returns a graph representation of an accept-symbol using
// Thompson's construction. Symbol wraps a rune and has no
// sub-expressions. Symbol essentially acts as a base case
// in the recursive Thompson() tree.
func (s Symbol) Thompson(counter int) (Alphabet, *Graph) {
	start := new(Vertex)
	end := new(Vertex)

	start.Id = counter
	end.Id = counter + 1

	start.Outgoing = []Edge{
		Edge{end, s}}

	return newAlphabet(s), &Graph{start, end}
}
