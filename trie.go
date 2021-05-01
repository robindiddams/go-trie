package trie

type node struct {
	next     map[rune]*node
	terminal bool
}

func newNode() *node {
	return &node{
		next: make(map[rune]*node),
	}
}

func newTerminalNode() *node {
	return &node{
		terminal: true,
		next:     make(map[rune]*node),
	}
}

// Trie is a trie for runes
type Trie struct {
	root *node
}

// NewTrie is the correct way to init a Trie
func NewTrie(values ...[]rune) Trie {
	t := Trie{
		root: newNode(),
	}
	for _, value := range values {
		t.Load(value)
	}
	return t
}

// NewTrieString is the same as NewTrie but accepts strings instead of []runes
func NewTrieString(values ...string) Trie {
	t := Trie{
		root: newNode(),
	}
	for _, value := range values {
		t.LoadString(value)
	}
	return t
}

// Load adds a []rune to the Trie
func (t *Trie) Load(value []rune) {
	if len(value) == 1 {
		t.root.next[value[0]] = newTerminalNode()
		return
	}
	current := t.root
	for i, r := range value {
		if i == len(value)-1 {
			// last element
			if current.next[r] == nil {
				current.next[r] = newTerminalNode()
			} else {
				// this may never happen
				current.next[r].terminal = true
			}
		} else {
			if current.next[r] == nil {
				current.next[r] = newNode()
			}
			current = current.next[r]
		}
	}
}

// LoadString is a convenience for t.Load([]rune(str))
func (t *Trie) LoadString(str string) {
	t.Load([]rune(str))
}

// SearchCallback is a function called when a match is found
// return false to keep searching
type SearchCallback func(hit []rune, at int) (done bool)

// SearchStringCallback is a function called when a match is found
// return false to keep searching
type SearchStringCallback func(hit string, at int) (done bool)

// SearchN searches s for any matches in the trie, returns the first n matches
func (t *Trie) SearchN(s []rune, n int) [][]rune {
	var found [][]rune
	if n < 1 {
		return found
	}
	t.SearchF(s, func(hit []rune, _ int) bool {
		found = append(found, hit)
		if len(found) >= n {
			return true
		}
		return false
	})
	return found
}

// SearchStringN searches s for any matches in the trie, returns the first n matches
func (t *Trie) SearchStringN(s string, n int) []string {
	var found []string
	if n < 1 {
		return found
	}
	t.SearchStringF(s, func(hit string, _ int) bool {
		found = append(found, hit)
		if len(found) >= n {
			return true
		}
		return false
	})
	return found
}

// Search searches s and returns any matches
func (t *Trie) Search(s []rune) [][]rune {
	var found [][]rune
	t.SearchF(s, func(hit []rune, _ int) bool {
		found = append(found, hit)
		return false
	})
	return found
}

// SearchString searches s and returns any matches
func (t *Trie) SearchString(s string) []string {
	var found []string
	t.SearchStringF(s, func(hit string, _ int) bool {
		found = append(found, hit)
		return false
	})
	return found
}

// SearchF searches input for trie matches and calls cb when it hits one,
// it will continue until cb returns true
func (t *Trie) SearchF(input []rune, cb SearchCallback) {
	var done bool
	for i := 0; i < len(input) && !done; i++ {
		r := input[i]
		if t.root.next[r] != nil {
			breadcrumbs := []rune{}
			var checkpoint []rune
			current := t.root
			for iter := i; iter < len(input); iter++ {
				// get the current rune
				newRune := input[iter]
				next := current.next[newRune]
				if next == nil {
					break
				}
				breadcrumbs = append(breadcrumbs, newRune)
				if next.terminal {
					checkpoint = breadcrumbs
				}
				current = next
			}
			if len(checkpoint) > 0 {
				done = cb(checkpoint, i)
				// advance our progress pass the emoji
				i += len(checkpoint) - 1
			}
		}
	}
}

// SearchStringF searches s for trie matches and calls cb when it hits one,
// it will continue until cb returns true
func (t *Trie) SearchStringF(s string, cb SearchStringCallback) {
	t.SearchF([]rune(s), func(hit []rune, at int) bool {
		return cb(string(hit), at)
	})
}
