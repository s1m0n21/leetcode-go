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
	var traversal func(*utils.TreeNode) []int

	traversal = func(node *utils.TreeNode) []int {
		if node == nil {
			return nil
		}

		values := make([]int, 0)
		values = append(values, traversal(node.Left)...)
		values = append(values, node.Val)
		values = append(values, traversal(node.Right)...)

		return values
	}

	return BSTIterator{traversal(root)}
}

func (this *BSTIterator) Next() int {
	num := this.nums[0]
	this.nums = this.nums[1:]
	return num
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}
