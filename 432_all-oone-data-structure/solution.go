/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/all-oone-data-structure/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/7 15:14
 */

package _all_oone_data_structure

import (
	"container/list"
)

type AllOne struct {
	nodes    map[string]*list.Element
	l        *list.List
}

type node struct {
	key   string
	count int
}

func Constructor() AllOne {
	return AllOne{
		nodes:  make(map[string]*list.Element),
		l:      list.New(),
	}
}

func (a *AllOne) Inc(key string) {
	if elem, in := a.nodes[key]; in {
		n := elem.Value.(node)
		n.count++
		elem.Value = n
		e := elem
		for e.Next() != nil && n.count >= e.Next().Value.(node).count {
			e = e.Next()
		}
		a.l.MoveAfter(elem, e)
	} else {
		n := node{
			key:   key,
			count: 1,
		}
		a.nodes[n.key] = a.l.PushFront(n)
	}
}

func (a *AllOne) Dec(key string) {
	elem := a.nodes[key]
	n := elem.Value.(node)
	n.count--
	elem.Value = n
	if n.count == 0 {
		a.l.Remove(elem)
		delete(a.nodes, n.key)
		return
	}
	e := elem
	for e.Prev() != nil && n.count <= e.Prev().Value.(node).count {
		e = e.Prev()
	}
	a.l.MoveBefore(elem, e)
}

func (a *AllOne) GetMaxKey() string {
	elem := a.l.Back()
	if elem == nil {
		return ""
	}
	return elem.Value.(node).key
}

func (a *AllOne) GetMinKey() string {
	elem := a.l.Front()
	if elem == nil {
		return ""
	}
	return elem.Value.(node).key
}
