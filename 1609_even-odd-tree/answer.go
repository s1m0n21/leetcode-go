/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/even-odd-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/20 3:41 下午
 */

package _even_odd_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
)

type elem struct {
	node  *utils.TreeNode
	level int
}

func isEvenOddTree(root *utils.TreeNode) bool {
	isEven := func(n int) bool { return n&1 == 0 }
	queue := []elem{{root, 0}}

	var last *elem
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]

		if last != nil && last.level != e.level {
			last = nil
		}

		if (isEven(e.level) && isEven(e.node.Val)) || (!isEven(e.level) && !isEven(e.node.Val)) {
			return false
		}
		if last != nil && ((isEven(e.level) && last.node.Val >= e.node.Val) || (!isEven(e.level) && last.node.Val <= e.node.Val)) {
			return false
		}

		if e.node.Left != nil {
			queue = append(queue, elem{e.node.Left, e.level + 1})
		}
		if e.node.Right != nil {
			queue = append(queue, elem{e.node.Right, e.level + 1})
		}

		last = &e
	}

	return true
}
