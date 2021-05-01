# go-trie

This is a trie in go. I use maps so it's not the most memory efficient, but it does the job. The purpose of this was mainly for parsing emojis out of text. Since the number of unicode codepoints for a single emoji is 1-n and some emojis are subsets of other emojis this makes parsing them for the purpose of extracting them very tricky and ideal for a trie.

PRs welcome (especially ones with emojis 🙌)!
