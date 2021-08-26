/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/design-hashset/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/23 9:13 上午
 */

package _design_hashset

type MyHashSet struct {
	m map[int]struct{}
}


func Constructor() MyHashSet {
	return MyHashSet{make(map[int]struct{})}
}


func (this *MyHashSet) Add(key int)  {
	if !this.Contains(key) {
		this.m[key] = struct{}{}
	}
}


func (this *MyHashSet) Remove(key int)  {
	if this.Contains(key) {
		delete(this.m, key)
	}
}


func (this *MyHashSet) Contains(key int) bool {
	_, has := this.m[key]
	return has
}
