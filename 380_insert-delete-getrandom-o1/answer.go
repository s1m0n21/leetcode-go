/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/insert-delete-getrandom-o1/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/6 11:23 AM
 */

package _insert_delete_getrandom_o1

import "math/rand"

type RandomizedSet struct {
	index  map[int]int
	values []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		index: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, has := this.index[val]; has {
		return false
	}
	this.values = append(this.values, val)
	this.index[val] = len(this.values) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, has := this.index[val]
	if !has {
		return false
	}
	this.values[len(this.values)-1], this.values[idx] = this.values[idx], this.values[len(this.values)-1]
	this.index[this.values[idx]] = idx
	this.values = this.values[:len(this.values)-1]
	delete(this.index, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}
