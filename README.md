# go-trie

This is a trie in go. I use maps so it's not the most memory efficient, but it does the job. The purpose of this was mainly for parsing emojis out of text. Since the number of unicode codepoints for a single emoji is 1-n and some emojis are subsets of other emojis this makes parsing them for the purpose of extracting them very tricky and ideal for a trie.

```Go
// create a trie with some string search values
trie := NewTrieString("foo", "bar", "foobar", "one")

// search the use the trie to search a string
results := trie.SearchString("there should be one foo, one bar, one foobar, and four 'one's in this string")
// results == [one foo one bar one foobar one]

// notice how it _doesn't_ match foo (or bar) in foobar, the trie will always match the maximum (ie. longest) search value. This is the magic of tries 🧙‍♂️
```

PRs welcome (especially ones with emojis 🙌)!
