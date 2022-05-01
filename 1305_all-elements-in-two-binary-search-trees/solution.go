/**
 * @Author         : s1m0n21
 * @Description    : Solotion of https://leetcode.cn/problems/all-elements-in-two-binary-search-trees/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/1 12:38
 */

package _all_elements_in_two_binary_search_trees

import "github.com/s1m0n21/leetcode-go/utils"

func getAllElements(root1 *utils.TreeNode, root2 *utils.TreeNode) []int {
	var ret []int

	nums1 := inorder(root1)
	nums2 := inorder(root2)

	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				ret = append(ret, nums1[i])
				i++
			} else {
				ret = append(ret, nums2[j])
				j++
			}
		} else if i < len(nums1) {
			ret = append(ret, nums1[i])
			i++
		} else {
			ret = append(ret, nums2[j])
			j++
		}
	}

	return ret
}

func inorder(root *utils.TreeNode) []int {
	var stack []*utils.TreeNode
	var ret []int

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}

	return ret
}
