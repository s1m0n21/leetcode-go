/**
 * @Author         : s1m0n21
 * @Description    : answer for https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/22 11:11 上午
 */

package _maximum_depth_of_n_ary_tre

import (
	"github.com/s1m0n21/leetcode-go/utils"
)

func maxDepth(root *utils.NTreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	for _, node := range root.Children {
		n := maxDepth(node)
		if n > ans {
			ans = n
		}
	}

	return ans + 1
}
