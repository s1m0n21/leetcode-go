/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/22 13:34
 */

package _average_of_levels_in_binary_tree

import "github.com/s1m0n21/leetcode-go/utils"

type element struct {
	level int
	node  *utils.TreeNode
}

func averageOfLevels(root *utils.TreeNode) []float64 {
	ans := []float64{float64(root.Val)}
	queue := []element{{0, root}}
	curr, cnt := -1, 1

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]

		if e.node.Left != nil {
			queue = append(queue, element{e.level + 1, e.node.Left})
		}
		if e.node.Right != nil {
			queue = append(queue, element{e.level + 1, e.node.Right})
		}

		if e.level != curr {
			if curr == -1 {
				curr = e.level
				continue
			}
			ans[curr] /= float64(cnt)
			ans = append(ans, float64(e.node.Val))
			curr = e.level
			cnt = 1
			continue
		}

		ans[curr] += float64(e.node.Val)
		cnt++
	}
	ans[curr] /= float64(cnt)

	return ans
}
