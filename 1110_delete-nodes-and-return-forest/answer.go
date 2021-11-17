/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/delete-nodes-and-return-forest/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/17 2:46 下午
 */

package _delete_nodes_and_return_forest

import "github.com/s1m0n21/leetcode-go/utils"

func delNodes(root *utils.TreeNode, to_delete []int) []*utils.TreeNode {
	var forest []*utils.TreeNode
	var dfs func(root *utils.TreeNode) *utils.TreeNode
	var deletes = make(map[int]struct{})

	for _, v := range to_delete {
		deletes[v] = struct{}{}
	}

	dfs = func(root *utils.TreeNode) *utils.TreeNode {
		if root == nil {
			return nil
		}

		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)

		if _, exist := deletes[root.Val]; exist {
			if root.Left != nil {
				forest = append(forest, root.Left)
			}
			if root.Right != nil {
				forest = append(forest, root.Right)
			}
			root = nil
		}

		return root
	}

	if root = dfs(root); root != nil {
		forest = append(forest, root)
	}

	return forest
}
