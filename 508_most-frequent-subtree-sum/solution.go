/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/most-frequent-subtree-sum/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/20 21:28
 */

package _most_frequent_subtree_sum

import "github.com/s1m0n21/leetcode-go/utils"

func findFrequentTreeSum(root *utils.TreeNode) []int {
	var dfs func(node *utils.TreeNode)
	var freq = make(map[int]int)
	var maxFreq int

	dfs = func(node *utils.TreeNode) {
		if node.Left == nil && node.Right == nil {
			goto DONE
		}

		if node.Left != nil {
			dfs(node.Left)
			node.Val += node.Left.Val
		}

		if node.Right != nil {
			dfs(node.Right)
			node.Val += node.Right.Val
		}

	DONE:
		freq[node.Val]++
		maxFreq = max(maxFreq, freq[node.Val])
	}

	dfs(root)

	var ret []int
	for n, f := range freq {
		if f == maxFreq {
			ret = append(ret, n)
		}
	}

	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
