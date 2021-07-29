/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-tree-paths/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/29 10:02 下午
 */

package _binary_tree_paths

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"strconv"
)

func binaryTreePaths(root *utils.TreeNode) []string {
	if root == nil {
		return nil
	}

	var res = make([]string, 0)

	if root.Left == nil && root.Right == nil {
		res =  append(res, strconv.Itoa(root.Val))
		return res
	}

	left := binaryTreePaths(root.Left)
	for i := 0; i < len(left); i++ {
		res = append(res, strconv.Itoa(root.Val) + "->" + left[i])
	}

	right := binaryTreePaths(root.Right)
	for i := 0; i < len(right); i++ {
		res = append(res, strconv.Itoa(root.Val) + "->" + right[i])
	}

	return res
}
