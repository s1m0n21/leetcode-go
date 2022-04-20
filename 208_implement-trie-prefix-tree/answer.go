/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/implement-trie-prefix-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/20 16:42
 */

package _implement_trie_prefix_tree

type trieNode struct {
	isWord   bool
	children [26]*trieNode
}

type Trie struct {
	root *trieNode
}

func Constructor() Trie {
	return Trie{&trieNode{}}
}

func (t *Trie) Insert(word string) {
	n, node := t.find(word)
	if n == len(word) {
		node.isWord = true
	} else {
		for i := n; i < len(word); i++ {
			node.children[word[i]-'a'] = &trieNode{}
			node = node.children[word[i]-'a']
		}
		node.isWord = true
	}
}

func (t *Trie) Search(word string) bool {
	n, node := t.find(word)
	return n == len(word) && node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	n, _ := t.find(prefix)
	return n == len(prefix)
}

func (t *Trie) find(s string) (int, *trieNode) {
	node := t.root
	for i, b := range s {
		if node.children[b-'a'] == nil {
			return i, node
		}
		node = node.children[b-'a']
	}
	return len(s), node
}
