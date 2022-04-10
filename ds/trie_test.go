package ds

import (
	"testing"
)

func TestTrie(t *testing.T) {
	myTrie := NewTrie()
	success := myTrie.Search("")
	if success {
		t.Error("Should not return true for empty string")
	}

	success = myTrie.Insert("")
	if success {
		t.Error("Should not be able to insert empty string")
	}

	success = myTrie.Insert("plant")
	if !success {
		t.Error("Should be able to insert lowercase word")
	}

	success = myTrie.Insert("ant")
	if !success {
		t.Error("Should be able to insert word that is a substring of previous word")
	}

	success = myTrie.Search("plan")
	if success {
		t.Error("Should not be able to find non-inserted word")
	}

	success = myTrie.Search("ant")
	if !success {
		t.Error("Should be able to find inserted word")
	}

	success = myTrie.Search("plant")
	if !success {
		t.Error("Should be able to find inserted word")
	}

	success = myTrie.Insert("25")
	if success {
		t.Error("Should not be able to insert non-alphabet characters")
	}

	success = myTrie.Insert("plant")
	if !success {
		t.Error("Should be able to insert duplicates")
	}
	myTrie.Insert("planche")
	myTrie.Insert("golang")
	myTrie.Insert("gopher")
	myTrie.Insert("zoo")
	allOptions := myTrie.WordOptions("")
	wantAllOptions := []string{"ant", "golang", "gopher", "planche", "plant", "zoo"}
	for i := 0; i < len(wantAllOptions); i++ {
		if allOptions[i] != wantAllOptions[i] {
			t.Errorf("Should be able to assemble all available options. Got: %+v, want: %+v", allOptions, wantAllOptions)
		}
	}

	onlyGoOptions := myTrie.WordOptions("go")
	wantGoOptions := []string{"golang", "gopher"}
	for i := 0; i < len(wantGoOptions); i++ {
		if onlyGoOptions[i] != wantGoOptions[i] {
			t.Errorf("Should be able to assemble all available options. Got: %+v, want: %+v", onlyGoOptions, wantGoOptions)
		}
	}
}
