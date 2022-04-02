/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/23 4:28 下午
 */

package _binary_tree_level_order_traversal

import "github.com/s1m0n21/leetcode-go/utils"

func levelOrder(root *utils.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	queue := []*utils.TreeNode{root}
	i := 0

	for len(queue) > 0 {
		q := queue
		queue = nil
		for len(q) > 0 {
			if len(ret) == i {
				ret = append(ret, []int{})
			}

			ret[i] = append(ret[i], q[0].Val)
			if q[0].Left != nil {
				queue = append(queue, q[0].Left)
			}
			if q[0].Right != nil {
				queue = append(queue, q[0].Right)
			}
			q = q[1:]
		}
		i++
	}

	return ret
}
