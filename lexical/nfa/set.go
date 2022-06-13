package nfa

// turns a Alphabet alphabet into a list of Symbols.
func (alpha Alphabet) Symbols() []Symbol {
	symbols := make([]Symbol, 0)

	for symbol := range alpha {
		symbols = append(symbols, symbol)
	}

	return symbols
}

// returns a new Alphabet containing the union of alpha and other.
func (alpha Alphabet) Union(other Alphabet) Alphabet {
	newAlpha := make(Alphabet)

	for symbol := range alpha {
		newAlpha[symbol] = true
	}

	for symbol := range other {
		newAlpha[symbol] = true
	}

	return newAlpha
}
