package ds

import (
	"regexp"
	"strings"
)

const alphabetLength = 26

// Make struct for node
type trieNode struct {
	children  [alphabetLength]*trieNode
	endOfWord bool
}

// Struct for trie
type trie struct {
	root *trieNode
}

func NewTrie() *trie {
	return &trie{&trieNode{}}
}

// Insert adds a new word to the trie
func (t *trie) Insert(word string) bool {
	if !isValid(word) {
		return false
	}
	word = strings.ToLower(word)
	currentNode := t.root

	for _, char := range word {
		alphabetIndex := char - 'a'
		if currentNode.children[alphabetIndex] == nil {
			currentNode.children[alphabetIndex] = &trieNode{}
		}
		currentNode = currentNode.children[alphabetIndex]
	}
	currentNode.endOfWord = true
	return true
}

// Search returns true if target word is contained in trie
func (t *trie) Search(word string) bool {
	if !isValid(word) {
		return false
	}
	finalNode, success := findNode(t.root, word)
	if !success {
		return false
	}
	return finalNode.endOfWord
}

func (t *trie) WordOptions(word string) []string {
	endOfTextNode, success := findNode(t.root, word)
	if !success {
		return []string{}
	}
	var wordOptions []string
	var generateWordOptions func(n *trieNode, index int, wordSoFar string) bool
	generateWordOptions = func(n *trieNode, index int, wordSoFar string) bool {
		if n == nil {
			return false
		}
		letter := string(rune(index) + 'a')
		newWord := wordSoFar + letter
		if n.endOfWord {
			wordOptions = append(wordOptions, newWord)
			var hasAChildNode bool
			for i := 0; i < alphabetLength; i++ {
				if n.children[i] != nil {
					hasAChildNode = true
				}
			}
			if !hasAChildNode {
				return true
			}
		}
		for i := 0; i < alphabetLength; i++ {
			generateWordOptions(n.children[i], i, newWord)
		}
		return true
	}

	for i := 0; i < alphabetLength; i++ {
		generateWordOptions(endOfTextNode.children[i], i, word)
	}
	return wordOptions
}

func findNode(starter *trieNode, currentText string) (*trieNode, bool) {
	currentNode := starter
	for _, char := range currentText {
		alphabetIndex := char - 'a'
		if currentNode.children[alphabetIndex] == nil {
			return nil, false
		}
		currentNode = currentNode.children[alphabetIndex]
	}
	return currentNode, true
}

func isValid(word string) bool {
	myRegex := regexp.MustCompile("[^a-zA-Z]")
	return !myRegex.MatchString(word) && len(word) > 0
}
