/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-tree-right-side-view/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/23 5:16 下午
 */

package _binary_tree_right_side_view

import "github.com/s1m0n21/leetcode-go/utils"

func rightSideView(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	var res = make([]int, 0)
	var queue = make([]*utils.TreeNode, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)
		res = append(res, queue[length-1].Val)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}
