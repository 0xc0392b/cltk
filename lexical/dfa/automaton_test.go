package dfa

import (
	"github.com/RHUL-CS-Projects/IndividualProject_2021_William.Santos/programming/cltk/lexical/nfa"
	"testing"
)

// constructs three simple test NFAs and runs a series of strings through
// them, checking whether they are accepted or rejected correctly.
func TestNFA(t *testing.T) {
	// ((a|b)c)*
	one := NewAutomaton(nfa.NewAutomaton(nfa.Closure{
		nfa.Concat{nfa.Union{nfa.Symbol('a'),
			nfa.Symbol('b')}, nfa.Symbol('c')}}))

	// (foo|bar|baz)*
	two := NewAutomaton(nfa.NewAutomaton(nfa.Closure{nfa.Union{
		nfa.Union{nfa.Concat{nfa.Concat{nfa.Symbol('f'),
			nfa.Symbol('o')}, nfa.Symbol('o')},
			nfa.Concat{nfa.Concat{nfa.Symbol('b'),
				nfa.Symbol('a')}, nfa.Symbol('r')}},
		nfa.Concat{nfa.Concat{nfa.Symbol('b'), nfa.Symbol('a')},
			nfa.Symbol('z')}}}))

	// hello (go|nfa)
	three := NewAutomaton(nfa.NewAutomaton(nfa.Concat{nfa.Concat{
		nfa.Concat{nfa.Concat{nfa.Concat{nfa.Concat{
			nfa.Symbol('h'), nfa.Symbol('e')},
			nfa.Symbol('l')}, nfa.Symbol('l')},
			nfa.Symbol('o')}, nfa.Symbol(' ')},
		nfa.Union{nfa.Concat{nfa.Symbol('g'), nfa.Symbol('o')},
			nfa.Concat{nfa.Concat{nfa.Symbol('n'),
				nfa.Symbol('f')}, nfa.Symbol('a')}}}))

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
