package opp

// takes a list of operator and operand symbols. returns an
// abstract symbol tree built by an operator-precedence parser.
func ParseTokens(tokens []Token) (*AST, error) {
	if postfix, err := shuntingYard(tokens); err != nil {
		return nil, err
	} else {
		ast := new(AST)
		ast.Root = postfixToAST(postfix)

		return ast, nil
	}
}
