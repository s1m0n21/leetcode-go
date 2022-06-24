/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/24 13:09
 */

package _find_largest_value_in_each_tree_row

import "github.com/s1m0n21/leetcode-go/utils"

func largestValues(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	q := []*utils.TreeNode{root}
	nextStart := root
	currLevel := -1
	findNext := false
	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nextStart {
			currLevel++
			ret = append(ret, node.Val)
			findNext = true
		} else {
			if node.Val > ret[currLevel] {
				ret[currLevel] = node.Val
			}
		}

		if node.Left != nil {
			q = append(q, node.Left)
			if findNext {
				nextStart = node.Left
				findNext = false
			}
		}

		if node.Right != nil {
			q = append(q, node.Right)
			if findNext {
				nextStart = node.Right
				findNext = false
			}
		}
	}

	return ret
}
