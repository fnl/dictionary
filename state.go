/* A linked state-node representation of the DFA. */
package lexicos

import (
	"fmt"
	"sort"
	"strings"
)

// A Lexicon can be implemented as a deterministic finite-state automaton (DFA).
// A DFA is a collection of states represented as vertices that are connected by outgoing edges, i.e., a directed (and, usually, acyclic) graph.
type graph struct {
	root state
}

// The edges are state transitions identified by ("labeled with") one symbol (rune) from the alphabet used to create the words in the Lexicon.
// In addition, the vertex stores a flag indicating if this is a final (but not necessarily terminal!) state in the graph.
// This final flag indicates if the particular path to this vertex (i.e., the sequence of symbols used as labels on the edges along this path) represents a word in the Lexicon.
type state struct {
	final       bool
	transitions map[rune]*state
}

// Make a (shallow) copy of a state.
func clone(vertex *state) (clone *state) {
	clone = new(state)
	clone.final = vertex.final
	if len(vertex.transitions) > 0 {
		clone.transitions = make(map[rune]*state)
		for k, v := range vertex.transitions {
			clone.transitions[k] = v
		}
	}
	return clone
}

// Create a fresh graph-based Lexicon.
func NewGraph() Lexicon {
	dfa := graph{state{}}
	dfa.root.transitions = make(map[rune]*state)
	return Lexicon(&dfa)
}

// Display the Lexicon using a regular (S-) expression-like syntax.
func (dfa *graph) String() string {
	return dfa.root.String()
}

// Display the DFA from this state on.
func (vertex *state) String() string {
	idx := len(vertex.transitions)
	if idx == 0 {
		return ""
	}
	str := make([]string, idx)
	idx = 0
	for k, v := range vertex.transitions {
		if v.final {
			str[idx] = fmt.Sprintf("%c$%s", k, v.String())
		} else {
			str[idx] = fmt.Sprintf("%c%s", k, v.String())
		}
		idx++
	}
	if idx == 1 {
		return str[0]
	}
	return fmt.Sprintf("(%s)", strings.Join(str, "|"))
}

// Set the final state or keep traversing.
// Creates new transition maps if necessary.
func (s *state) setState(word []rune, idx int) bool {
	idx++
	if len(word) > idx {
		if len(s.transitions) == 0 {
			s.transitions = make(map[rune]*state)
		}
		return s.addWord(word, idx)
	} else {
		existed := s.final
		s.final = true
		return !existed
	}
}

// Add the word path starting at idx to this state.
func (s *state) addWord(word []rune, idx int) bool {
	if succ, exists := s.transitions[word[idx]]; exists {
		return succ.setState(word, idx)
	} else {
		succ = new(state)
		s.transitions[word[idx]] = succ
		return succ.setState(word, idx)
	}
}

// Return the state and offset in word that represents the longest common prefix in the DFA.
func (dfa *graph) commonPrefix(word []rune, idx int) (*state, int) {
	s := &dfa.root
	for len(s.transitions) > 0 && len(word) > idx {
		if next, exists := s.transitions[word[idx]]; exists {
			s = next
			idx++
		} else {
			break
		}
	}
	return s, idx
}

func (dfa *graph) replaceOrRegister(s *state) {
	t := len(s.transitions)
	keys := make([]int, t)
	i := 0
	for k, _ := range s.transitions {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)
	s = s.transitions[rune(keys[t-1])]
	if len(s.transitions) > 0 {
		dfa.replaceOrRegister(s)
	}
	// TODO...
}

// Minimize the DFA at this state.
func cleanup(s *state) {
}

func (dfa *graph) Insert(word string) bool {
	symbols := []rune(word)
	s, i := dfa.commonPrefix(symbols, 0)
	if len(s.transitions) > 0 {
		dfa.replaceOrRegister(s)
	}
	if len(symbols) > i {
		return s.addWord(symbols, i)
	} else {
		return false
	}
}

func (dfa *graph) Delete(word string) bool {
	return false
}

func (dfa *graph) Contains(word string) bool {
	return false
}

func (dfa *graph) Words(phrase string, offset int) []string {
	return nil
}

func (dfa *graph) Lookup(prefix string) []string {
	return nil
}
