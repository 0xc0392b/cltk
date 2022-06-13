package nfa

import "testing"

// constructs three simple test NFAs and runs a series of strings through
// them, checking whether they are accepted or rejected correctly.
func TestNFA(t *testing.T) {
	// ((a|b)c)*
	one := NewAutomaton(Closure{Concat{Union{Symbol('a'),
		Symbol('b')}, Symbol('c')}})

	// (foo|bar|baz)*
	two := NewAutomaton(Closure{Union{Union{
		Concat{Concat{Symbol('f'), Symbol('o')}, Symbol('o')},
		Concat{Concat{Symbol('b'), Symbol('a')}, Symbol('r')}},
		Concat{Concat{Symbol('b'), Symbol('a')}, Symbol('z')}}})

	// hello (go|nfa)
	three := NewAutomaton(Concat{Concat{Concat{Concat{Concat{Concat{
		Symbol('h'), Symbol('e')}, Symbol('l')},
		Symbol('l')}, Symbol('o')}, Symbol(' ')},
		Union{Concat{Symbol('g'), Symbol('o')},
			Concat{Concat{Symbol('n'), Symbol('f')}, Symbol('a')}}})

	for _, test := range []struct {
		automaton *Automaton
		input     string
		expected  bool
	}{
		// test first NFA
		{one, "", true},
		{one, "ac", true},
		{one, "bc", true},
		{one, "acbcac", true},
		{one, "acacacbcbcbc", true},
		{one, "bcbcbcacacac", true},
		{one, "bcbcbcacacacd", false},
		{one, "a", false},
		{one, "b", false},
		{one, "c", false},
		{one, "aabbcc", false},
		{one, "abcabcabc", false},

		// test second NFA
		{two, "", true},
		{two, "foo", true},
		{two, "bar", true},
		{two, "baz", true},
		{two, "foobarbaz", true},
		{two, "foofoofoobazbazbaz", true},
		{two, "foobarfoobarbaz", true},
		{two, "foobarandbaz", false},
		{two, "foo bar baz", false},

		// test third NFA
		{three, "hello go", true},
		{three, "hello nfa", true},
		{three, "", false},
		{three, "hello", false},
		{three, "go", false},
		{three, "nfa", false},
		{three, "hello world", false},
		{three, "hellogo", false},
		{three, "hellonfa", false},
	} {
		if accepts := test.automaton.AcceptsStr(
			test.input); accepts != test.expected {
			t.Error(test.input, accepts,
				"expected", test.expected)
		}
	}
}
