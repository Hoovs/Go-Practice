package trees

type trieNode struct {
	links map[rune]*trieNode
	isEnd bool
}

type Trie struct {
	root *trieNode
}

func InitTrie() *Trie {
	return &Trie{root: &trieNode{links: make(map[rune]*trieNode, 0), isEnd: false}}
}

func (t *Trie) FindPrefix(prefix string) bool {
	root := t.root
	for _, c := range prefix {
		if v, ok := root.links[c]; ok {
			root = v
		} else {
			return false
		}
	}

	return true
}

func (t *Trie) Contains(word string) bool {
	root := t.root
	for _, c := range word {
		if v, ok := root.links[c]; ok {
			root = v
		} else {
			return false
		}
	}

	return root.isEnd
}

func (t *Trie) Insert(word string) {
	curNode := t.root
	for _, c := range word {
		if _, ok := curNode.links[c]; !ok {
			curNode.links[c] = &trieNode{links: make(map[rune]*trieNode, 0), isEnd: false}
		}
		curNode = curNode.links[c]
	}

	curNode.isEnd = true
}
