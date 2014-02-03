/* A hash-map based representation of a Lexicon's DFA. */
package lexicos

// This representation is based on a hash-map implementation of the directed graph.
// In this case, as compared to the state-based representation, only final states are "stored" internally.
// The graph is stored in a hash-map that uses the concatenation of current state and label as keys and the target states as values.
// While this implementation is faster and requires less memmory for all but tiny DFAs, it has an upper boundary defined as the number of states it can store (uint32, 2^32, or about 4 billion nodes).
type Trie struct {
	stateCount  uint32
	transitions map[uint64]uint32
	finalStates map[uint32]bool
}

// Create a new hash-map based representation.
func New() *Trie {
	t := new(Trie)
	t.stateCount = 0
	t.transitions = make(map[uint64]uint32)
	t.finalStates = make(map[uint32]bool)
	return t
}

// Concatenate the current state and the transition's label to a key.
func makeKey(state uint32, label rune) uint64 {
	key := uint64(state)
	key <<= 32
	key += uint64(label)
	return key
}

// Fetch a new state.
// Panics if there are no more states (max. states is 2^32).
func (t *Trie) Create() (state uint32) {
	state = t.stateCount
	t.stateCount++
	if 0 == t.stateCount {
		panic("ran out of states")
	}
	return state
}

// Add a labeled transition from a source to a target state to the graph.
func (t *Trie) Add(srcState uint32, label rune, dstState uint32) {
	t.transitions[makeKey(srcState, label)] = dstState
}

// Walk to the next state along the correctly labeled vertex.
// If the secondary value is false, the walk failed.
func (t *Trie) Walk(srcState uint32, label rune) (dstState uint32, ok bool) {
	dstState, ok = t.transitions[makeKey(srcState, label)]
	return
}
