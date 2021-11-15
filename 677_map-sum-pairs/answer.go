/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/map-sum-pairs/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/15 9:32 上午
 */

package _map_sum_pairs

type TrieNode struct {
	next [26]*TrieNode
	val  int
}

type MapSum struct {
	trie *TrieNode
}

func Constructor() MapSum {
	return MapSum{&TrieNode{}}
}

func (this *MapSum) Insert(key string, val int) {
	var tail = this.trie
	for _, b := range key {
		i := b - 'a'
		if tail.next[i] == nil {
			node := &TrieNode{}
			tail.next[i] = node
		}
		tail = tail.next[i]
	}
	tail.val = val
}

func (this *MapSum) Sum(prefix string) int {
	var tail = this.trie
	for _, b := range prefix {
		i := b - 'a'
		if tail.next[i] == nil {
			return 0
		}
		tail = tail.next[i]
	}

	var sum int
	var stack []*TrieNode
	stack = append(stack, tail)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += node.val
		for i := 25; i >= 0; i-- {
			if node.next[i] != nil {
				stack = append(stack, node.next[i])
			}
		}
	}

	return sum
}
