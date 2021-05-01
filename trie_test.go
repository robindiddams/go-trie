package trie

import (
	"testing"
)

func assertLenRuneArr(t *testing.T, slice [][]rune, length int) {
	if len(slice) != length {
		t.Fatal("slice", slice, "should have length", length, "but has length", len(slice))
	}
}

func assertLenString(t *testing.T, slice []string, length int) {
	if len(slice) != length {
		t.Fatal("slice", slice, "should have length", length, "but has length", len(slice))
	}
}

func assertEqual(t *testing.T, a []rune, b []rune) {
	if string(a) != string(b) {
		t.Fatal("values are not equal:", string(a), string(b))
	}
}

func assertEqualString(t *testing.T, a string, b string) {
	if a != b {
		t.Fatal("values are not equal:", string(a), string(b))
	}
}

func TestLoad(t *testing.T) {
	trie := NewTrie([]rune{'a', 'b'})
	matches := trie.SearchString("abc")
	assertLenString(t, matches, 1)
}

func TestLoadString(t *testing.T) {
	trie := NewTrieString("ab")
	matches := trie.SearchString("abc")
	assertLenString(t, matches, 1)
}

func TestTrie(t *testing.T) {
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie := NewTrie(thermometer, violin)
	emojis := trie.SearchStringN("I love to play the "+string(thermometer)+" while i eat a "+string(violin), 2)
	assertLenString(t, emojis, 2)
}

func TestSearch(t *testing.T) {
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie := NewTrie(thermometer, violin)
	emojis := trie.Search([]rune("I love to play the " + string(thermometer) + " while i eat a " + string(violin)))
	assertLenRuneArr(t, emojis, 2)
	assertEqual(t, emojis[0], thermometer)
	assertEqual(t, emojis[1], violin)
}

func TestSearchN(t *testing.T) {
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie := NewTrie(thermometer, violin)
	emojis := trie.SearchN([]rune("I love to play the "+string(thermometer)+"while i eat a "+string(violin)), 1)
	assertLenRuneArr(t, emojis, 1)
	emojis = trie.SearchN([]rune("I love to play the "+string(thermometer)+"while i eat a "+string(violin)), 2)
	assertLenRuneArr(t, emojis, 2)
	emojis = trie.SearchN([]rune("I love to play the "+string(thermometer)+"while i eat a "+string(violin)), 0)
	assertLenRuneArr(t, emojis, 0)
}

func TestSearchStringN(t *testing.T) {
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie := NewTrie(thermometer, violin)
	emojis := trie.SearchStringN("I love to play the "+string(thermometer)+"while i eat a "+string(violin), 1)
	assertLenString(t, emojis, 1)
	emojis = trie.SearchStringN("I love to play the "+string(thermometer)+"while i eat a "+string(violin), 2)
	assertLenString(t, emojis, 2)
}

func TestSearchStringNZero(t *testing.T) {
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie := NewTrie(thermometer, violin)
	emojis := trie.SearchStringN("I love to play the "+string(thermometer)+"while i eat a "+string(violin), 0)
	assertLenString(t, emojis, 0)
	emojis = trie.SearchStringN("I love to play the "+string(thermometer)+"while i eat a "+string(violin), -100)
	assertLenString(t, emojis, 0)
}

func TestZWJEmoji(t *testing.T) {
	femaleSign := []rune{0x2640, 0xfe0f}
	// 129335, 8205, 9792, 65039
	womanShrugging := []rune{0x1f937, 0x200d, 0x2640, 0xfe0f}
	trie := NewTrie(womanShrugging, femaleSign)
	var emojis [][]rune
	trie.SearchF([]rune("fix tests!! or did I "+string(womanShrugging)), func(hit []rune, _ int) bool {
		emojis = append(emojis, hit)
		return false
	})
	assertLenRuneArr(t, emojis, 1)
	assertEqual(t, womanShrugging, emojis[0])
}
