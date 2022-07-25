/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/complete-binary-tree-inserter/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/25 09:41
 */

package _complete_binary_tree_inserter

import "github.com/s1m0n21/leetcode-go/utils"

type CBTInserter struct {
	root  *utils.TreeNode
	queue []*utils.TreeNode
}

func Constructor(root *utils.TreeNode) CBTInserter {
	queue := []*utils.TreeNode{root}

	node := root
	for node.Left != nil && node.Right != nil {
		queue = queue[1:]
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
		node = queue[0]
	}
	if node.Left != nil {
		queue = append(queue, node.Left)
	}

	return CBTInserter{root, queue}
}

func (c *CBTInserter) Insert(val int) int {
	node := c.queue[0]
	newNode := &utils.TreeNode{Val: val}
	if node.Left == nil {
		node.Left = newNode
	} else {
		node.Right = newNode
	}
	c.queue = append(c.queue, newNode)

	if node.Left != nil && node.Right != nil {
		c.queue = c.queue[1:]
	}

	return node.Val
}

func (c *CBTInserter) Get_root() *utils.TreeNode {
	return c.root
}
