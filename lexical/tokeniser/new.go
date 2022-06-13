package tokeniser

import "github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/regex"

// constructs and returns a new tokeniser from a list of token
// definitions.
func New(tokenDefs ...TokenDef) (*Tokeniser, error) {
	tokeniser := new(Tokeniser)
	tokeniser.machines = make([]machine, len(tokenDefs))

	// create a regex engine for each token's pattern and add
	// it to the tokeniser as a string accepting "machine".
	for i, tokenDef := range tokenDefs {
		if r, err := regex.New(tokenDef.Pattern); err != nil {
			return nil, err
		} else {
			tokeniser.machines[i] = machine{
				name:    tokenDef.Name,
				grammar: r,
			}
		}
	}

	return tokeniser, nil
}
