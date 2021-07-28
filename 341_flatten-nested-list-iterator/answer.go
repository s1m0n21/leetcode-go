/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/flatten-nested-list-iterator/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/23 3:18 下午
 */

package _flatten_nested_list_iterator

import "github.com/s1m0n21/leetcode-go/utils"

type NestedIterator struct {
	n []int
}

func Constructor(nestedList []*utils.NestedInteger) *NestedIterator {
	return &NestedIterator{
		n: flatten(nestedList),
	}
}

func flatten(nestedList []*utils.NestedInteger) []int {
	var values = make([]int, 0)
	for _, i := range nestedList {
		if i == nil {
			continue
		}

		if i.IsInteger() {
			values = append(values, i.GetInteger())
		} else {
			values = append(values, flatten(i.GetList())...)
		}
	}

	return values
}

func (this *NestedIterator) Next() int {
	var n int
	if this.HasNext() {
		n = this.n[0]
		this.n = this.n[1:]
	}
	return n
}

func (this *NestedIterator) HasNext() bool {
	if len(this.n) > 0 {
		return true
	}
	return false
}
