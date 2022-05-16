/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/sort-array-by-parity-ii/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/16 10:36
 */

package _sort_array_by_parity_ii

func sortArrayByParityII(nums []int) []int {
	for i, j := 0, 1; i < len(nums); i += 2 {
		if nums[i]%2 == 1 {
			for nums[j]%2 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
