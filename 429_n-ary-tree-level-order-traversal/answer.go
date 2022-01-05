/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/5 1:10 PM
 */

package _n_ary_tree_level_order_traversal

import "github.com/s1m0n21/leetcode-go/utils"

type element struct {
	node  *utils.NTreeNode
	level int
}

func levelOrder(root *utils.NTreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	var queue []element

	queue = append(queue, element{root, 0})
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]

		if e.level >= len(ans) {
			ans = append(ans, []int{e.node.Val})
		} else {
			ans[e.level] = append(ans[e.level], e.node.Val)
		}

		for _, node := range e.node.Children {
			queue = append(queue, element{node, e.level + 1})
		}
	}

	return ans
}
