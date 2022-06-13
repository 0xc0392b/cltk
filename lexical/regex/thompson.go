package regex

import (
	"github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"
	"github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/opp"
)

// rewrites an abstract syntax tree from the opp package into
// a type that satisfies nfa.Thompsoner.
func astToThompson(node *opp.ASTNode) nfa.Thompsoner {
	token := node.Token
	lexeme := token.Lexeme()
	symbol := lexeme[0]

	switch symbol {
	// recursively build left and right subtrees for the union
	// operator.
	case '|':
		lExp := astToThompson(node.Children[0])
		rExp := astToThompson(node.Children[1])
		return nfa.Union{lExp, rExp}

	// recursively build left and right subtrees for the concat
	// operator.
	case '&':
		lExp := astToThompson(node.Children[0])
		rExp := astToThompson(node.Children[1])
		return nfa.Concat{lExp, rExp}

	// recursively build subtree for the closure operator
	case '*':
		exp := astToThompson(node.Children[0])
		return nfa.Closure{exp}

	// base case: accept the symbol
	default:
		return nfa.Symbol(symbol)
	}
}
