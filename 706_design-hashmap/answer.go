/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/design-hashmap/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/9 9:20 上午
 */

package _design_hashmap

import "container/list"

const base = 1009

type MyHashMap struct {
	data []list.List
}

type entry struct {
	key, val int
}

func Constructor() MyHashMap {
	return MyHashMap{data: make([]list.List, base)}
}

func (m *MyHashMap) Put(key int, value int) {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap) Get(key int) int {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.val
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m.data[h].Remove(e)
		}
	}
}

func hash(key int) int {
	return key % base
}
