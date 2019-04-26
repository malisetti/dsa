package trie

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestMakeTrie(t *testing.T) {
	const dictPath = "/usr/share/dict/words"
	f, _ := os.Open(dictPath)
	r := bufio.NewReader(f)

	words := make(chan string)
	go func() {
		for {
			word, err := r.ReadString('\n')
			if err == io.EOF {
				break
			}

			if len(word) == 0 {
				t.Fail()
			}

			words <- word
		}
		close(words)
	}()

	root := MakeTrie(words)

	wordsToSearch := []string{
		"zits",
		"zoom",
		"zucchini",
		"yoghourts",
		"études",
		"étude",
	}

	for _, word := range wordsToSearch {
		if nil == Search(root, word) {
			t.Fail()
		}
	}

	Prefix(root, "sun")
}
