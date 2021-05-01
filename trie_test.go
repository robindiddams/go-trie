package trie

import (
	"testing"
)

func assertLen(t *testing.T, slice [][]rune, length int) {
	if len(slice) != length {
		t.Fatal("slice", slice, "should have length", length, "but has length", len(slice))
	}
}
func assertEqual(t *testing.T, a []rune, b []rune) {
	if string(a) != string(b) {
		t.Fatal("values are not equal:", string(a), string(b))
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie.Load(thermometer)
	trie.Load(violin)
	emojis := trie.Search("I love to play the "+string(thermometer)+" while i eat a "+string(violin), 2)
	assertLen(t, emojis, 2)
}

func TestTrieLimit(t *testing.T) {
	trie := NewTrie()
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie.Load(thermometer)
	trie.Load(violin)
	emojis := trie.Search("I love to play the "+string(thermometer)+"while i eat a "+string(violin), 1)
	assertLen(t, emojis, 1)
	emojis = trie.Search("I love to play the "+string(thermometer)+"while i eat a "+string(violin), -1)
	assertLen(t, emojis, 2)
}

func TestZWJEmoji(t *testing.T) {
	trie := NewTrie()
	femaleSign := []rune{0x2640, 0xfe0f}
	// 129335, 8205, 9792, 65039
	womanShrugging := []rune{0x1f937, 0x200d, 0x2640, 0xfe0f}
	trie.Load(womanShrugging)
	trie.Load(femaleSign)
	var emojis [][]rune
	trie.SearchF("fix tests!! or did I "+string(womanShrugging), func(hit []rune, _ int) bool {
		emojis = append(emojis, hit)
		return false
	})
	assertLen(t, emojis, 1)
	assertEqual(t, womanShrugging, emojis[0])
}
