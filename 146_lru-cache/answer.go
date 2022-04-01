/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/lru-cache/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/1 10:43
 */

package _lru_cache

type LRUCache struct {
	front  *linkedList
	back   *linkedList
	record map[int]*linkedList
	cap    int
	size   int
}

type linkedList struct {
	key, value int
	prev, next *linkedList
}

func Constructor(capacity int) LRUCache {
	back := &linkedList{}
	front := &linkedList{next: back}
	back.prev = front

	return LRUCache{
		front:  front,
		back:   back,
		record: make(map[int]*linkedList),
		cap:    capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	node, has := c.record[key]
	if !has {
		return -1
	}
	c.moveToFront(node)
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	if node, has := c.record[key]; has {
		node.value = value
		c.moveToFront(node)
		return
	}

	if c.size == c.cap {
		c.remove(c.back.prev)
	}

	c.record[key] = &linkedList{key: key, value: value}
	c.insertToFront(c.record[key])
}

func (c *LRUCache) moveToFront(node *linkedList) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = c.front.next
	c.front.next.prev = node
	c.front.next = node
	node.prev = c.front
}

func (c *LRUCache) insertToFront(node *linkedList) {
	c.front.next.prev = node
	node.next = c.front.next
	c.front.next = node
	node.prev = c.front
	c.size++
}

func (c *LRUCache) remove(node *linkedList) {
	node.prev.next = node.next
	node.next.prev = node.prev
	c.size--
	delete(c.record, node.key)
}
