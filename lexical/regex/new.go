package regex

import (
	"github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/dfa"
	"github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"
	"github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/opp"
)

// creates and returns a new regex engine from a pattern string.
// the pattern is parsed and transformed into an abstract syntax
// tree using an operator precedence parser. the AST is used to
// build the NFA automaton which is then used to build the DFA
// automaton.
func New(pattern string) (*Regex, error) {
	expression := strToExpression(pattern)

	if ast, err := opp.ParseTokens(expression); err != nil {
		return nil, err
	} else {
		// pattern -> ast -> thompson's construction -> NFA -> DFA
		thompson := astToThompson(ast.Root)
		nfaAutomaton := nfa.NewAutomaton(thompson)
		dfaAutomaton := dfa.NewAutomaton(nfaAutomaton)

		r := new(Regex)
		r.grammar = dfaAutomaton

		return r, nil
	}
}
