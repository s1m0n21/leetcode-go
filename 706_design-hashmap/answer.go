/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/design-hashmap/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/9 9:20 上午
 */

package _design_hashmap

type MyHashMap struct {
	m map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make(map[int]int)}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	this.m[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if val, has := this.m[key]; has {
		return val
	}

	return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	delete(this.m, key)
}
