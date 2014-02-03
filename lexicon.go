package lexicos

// A Lexicon is a data structure that hold a collection of words (byte sequences).
type Lexicon interface {
	// Add a word to the Lexicon.
	// Return false if the word existed in the Lexicon.
	Insert(word string) bool
	// Remove a word from the Lexicon.
	// Return true if the word existed in the Lexicon.
	Delete(word string) bool
	// Check if the word can be found in the lexicon.
	Contains(word string) bool
	// Return all words in the Lexicon that match at that offset in the phrase.
	Words(phrase string, offset int) []string
	// Return all words that begin with this prefix.
	Lookup(prefix string) []string
	String() string
}
