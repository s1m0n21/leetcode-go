/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/lfu-cache/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/15 09:33
 */

package _lfu_cache

import (
	"container/list"
)

type LFUCache struct {
	record  map[int]*list.Element
	freq    map[int]*list.List
	cap     int
	size    int
	minFreq int
}

type node struct {
	key, val, freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		record: make(map[int]*list.Element),
		freq:   make(map[int]*list.List),
		cap:    capacity,
	}
}

func (c *LFUCache) Get(key int) int {
	if elem, has := c.record[key]; has {
		n := elem.Value.(node)
		c.addFreq(n, elem, n.val)
		return n.val
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if c.cap == 0 {
		return
	}

	if elem, has := c.record[key]; has {
		n := elem.Value.(node)
		c.addFreq(n, elem, value)
	} else {
		n := node{
			key:  key,
			val:  value,
			freq: 1,
		}
		c.size++
		if c.size > c.cap {
			c.deleteNode()
		}
		c.minFreq = 1
		c.addNode(n)
	}
}

func (c *LFUCache) deleteNode() {
	elem := c.freq[c.minFreq].Back()
	n := elem.Value.(node)
	c.freq[n.freq].Remove(elem)
	delete(c.record, n.key)
	if c.freq[n.freq].Len() == 0 {
		delete(c.freq, n.freq)
	}
	c.size--
}

func (c *LFUCache) addFreq(n node, elem *list.Element, val int) {
	c.freq[n.freq].Remove(elem)
	if c.freq[n.freq].Len() == 0 {
		delete(c.freq, n.freq)
		if n.freq == c.minFreq {
			c.minFreq++
		}
	}
	n.freq++
	n.val = val
	c.addNode(n)
}

func (c *LFUCache) addNode(n node) {
	var e *list.Element
	if l, has := c.freq[n.freq]; has {
		e = l.PushFront(n)
	} else {
		c.freq[n.freq] = list.New()
		e = c.freq[n.freq].PushFront(n)
	}
	c.record[n.key] = e
}
