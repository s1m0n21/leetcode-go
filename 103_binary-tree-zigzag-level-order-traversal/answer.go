/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/23 5:06 下午
 */

package _binary_tree_zigzag_level_order_traversal

import "github.com/s1m0n21/leetcode-go/utils"

type elem struct {
	node  *utils.TreeNode
	level int
}

func zigzagLevelOrder(root *utils.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res = make([][]int, 0)
	var queue = make([]elem, 0)

	queue = append(queue, elem{
		node:  root,
		level: 0,
	})

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]

		if e.level == len(res) {
			res = append(res, []int{})
		}
		res[e.level] = append(res[e.level], e.node.Val)

		if e.node.Left != nil {
			queue = append(queue, elem{
				node:  e.node.Left,
				level: e.level + 1,
			})
		}

		if e.node.Right != nil {
			queue = append(queue, elem{
				node:  e.node.Right,
				level: e.level + 1,
			})
		}
	}

	for i, n := range res {
		if i % 2 != 0 {
			for j := 0; j < len(n)/2; j++ {
				n[j], n[len(n)-j-1] = n[len(n)-j-1], n[j]
			}
		}
	}

	return res
}
