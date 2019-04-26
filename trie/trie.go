package trie

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

// Node is a part of tree
type Node struct {
	Descendants map[rune]*Node
}

// MakeTrie builds a trie out of words
func MakeTrie(words <-chan string) *Node {
	root := &Node{
		Descendants: make(map[rune]*Node),
	}
	for word := range words {
		PutWordInTrie(root, word)
	}

	return root
}

// PutWordInTrie puts a word in trie
func PutWordInTrie(root *Node, word string) {
	current := root
	for _, char := range word {
		if current.Descendants == nil {
			current.Descendants = make(map[rune]*Node)
		}

		if node, ok := current.Descendants[char]; ok {
			current = node
		} else {
			current.Descendants[char] = &Node{
				Descendants: make(map[rune]*Node),
			}
			current = current.Descendants[char]
		}
	}
}

// Search searches if word exists in the trie of root
func Search(root *Node, word string) *Node {
	current := root
	for _, char := range word {
		if d, ok := current.Descendants[char]; ok {
			current = d
		} else {
			return nil
		}
	}
	return current
}

// Prefix finds the words that have prefix
func Prefix(root *Node, prefix string) error {
	current := Search(root, prefix)
	var word []rune
	PrintWords(current, prefix, word)

	return nil
}

// PrintWords prints words from a node
func PrintWords(root *Node, prefix string, word []rune) {
	if len(root.Descendants) == 0 {
		fmt.Printf("%s%s", prefix, string(word))
	} else {
		keys := make([]string, 0, len(root.Descendants))
		for c := range root.Descendants {
			keys = append(keys, string(c))
		}

		sort.Strings(keys)

		for _, s := range keys {
			c, _ := utf8.DecodeRune([]byte(s))
			d := root.Descendants[c]
			PrintWords(d, prefix, append(word, c))
		}
	}
}
