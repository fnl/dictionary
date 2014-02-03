package lexicos

import "testing"

func TestNewGraph(t *testing.T) {
	n := NewGraph()
	print(Lexicon(n))
}

func TestString(t *testing.T) {
	n := NewGraph()
	if n.String() != "" {
		t.Errorf("default should be an empty string, got %q", n.String())
	}
	n.Insert("x")
	n.Insert("yz")
	if n.String() != "(x$|yz$)" {
		t.Errorf("expected '(x$|yz$)', got %q", n.String())
	}
}

func TestInsert(t *testing.T) {
	words := [8]string{"heard", "here", "herd", "head", "hard", "her", "had", "he"}
	graph := NewGraph()
	for i := range words {
		if !graph.Insert(words[i]) {
			t.Errorf("received false on Insert of new word %q", words[i])
		}
	}
	if graph.Insert(words[0]) {
		t.Errorf("received true on Insert of existing word %q", words[0])
	}
	if graph.String() != "h(e$(a(rd$|d$)|r$(e$|d$))|a(rd$|d$))" {
		t.Errorf("expected 'h(e$(a(rd$|d$)|r$(e$|d$))|a(rd$|d$))', got %q", graph.String())
	}
}

func TestUnicode(t *testing.T) {
	words := [...]string{"The", "quick", "brown", "狐", "jumped", "over", "the", "lazy", "犬"}
	graph := NewGraph()
	for i := range words {
		if !graph.Insert(words[i]) {
			t.Errorf("received false on Insert of new word %q", words[i])
		}
	}
	// TODO: test other methods...
}
