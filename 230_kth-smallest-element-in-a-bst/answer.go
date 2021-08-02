/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/2 4:16 上午
 */

package _kth_smallest_element_in_a_bst

import "github.com/s1m0n21/leetcode-go/utils"

func kthSmallest(root *utils.TreeNode, k int) int {
	return travel(root)[k-1]
}

func travel(node *utils.TreeNode) []int {
	if node == nil {
		return nil
	}

	values := make([]int, 0)
	values = append(values, travel(node.Left)...)
	values = append(values, node.Val)
	values = append(values, travel(node.Right)...)

	return values
}
