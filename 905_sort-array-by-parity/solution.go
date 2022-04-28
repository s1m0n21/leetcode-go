/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/sort-array-by-parity/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/4/28 11:17
 */

package _sort_array_by_parity

func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]%2 == 0 {
			l++
		}
		for l < r && nums[r]%2 != 0 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return nums
}
