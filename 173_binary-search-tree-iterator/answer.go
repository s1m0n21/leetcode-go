/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-search-tree-iterator/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/8 8:19 下午
 */

package _binary_search_tree_iterator

import "github.com/s1m0n21/leetcode-go/utils"

type BSTIterator struct {
	nums []int
}

func Constructor(root *utils.TreeNode) BSTIterator {
	var stack []*utils.TreeNode
	var nums []int

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}

	return BSTIterator{nums}
}

func (this *BSTIterator) Next() int {
	num := this.nums[0]
	this.nums = this.nums[1:]
	return num
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}
