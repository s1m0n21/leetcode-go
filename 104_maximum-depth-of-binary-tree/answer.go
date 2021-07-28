/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/26 11:56 上午
 */

package _maximum_depth_of_binary_tree

import "github.com/s1m0n21/leetcode-go/utils"

func maxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	var queue = make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	depth := 0

	for len(queue) > 0 {
		length := len(queue)
		depth += 1
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

	return depth
}
