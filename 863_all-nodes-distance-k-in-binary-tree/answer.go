/**
 * @Author         : s1m0n21
 * @Description    : Ans of https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/16 11:58 AM
 */

package _all_nodes_distance_k_in_binary_tree

import "github.com/s1m0n21/leetcode-go/utils"

func distanceK(root *utils.TreeNode, target *utils.TreeNode, k int) []int {
	var ans []int
	parents := findParent(root)

	var dfs func(prev, curr *utils.TreeNode, depth int)
	dfs = func(prev, curr *utils.TreeNode, depth int) {
		if curr == nil {
			return
		}

		if depth == k {
			ans = append(ans, curr.Val)
			return
		}

		if curr.Left != prev {
			dfs(curr, curr.Left, depth+1)
		}

		if curr.Right != prev {
			dfs(curr, curr.Right, depth+1)
		}

		if parents[curr] != prev {
			dfs(curr, parents[curr], depth+1)
		}
	}

	dfs(nil, target, 0)

	return ans
}

func findParent(node *utils.TreeNode) map[*utils.TreeNode]*utils.TreeNode {
	var parents = make(map[*utils.TreeNode]*utils.TreeNode)
	var dfs func(node, parent *utils.TreeNode)

	dfs = func(node, parent *utils.TreeNode) {
		if node == nil {
			return
		}
		parents[node] = parent
		dfs(node.Left, node)
		dfs(node.Right, node)
	}

	dfs(node, nil)

	return parents
}
