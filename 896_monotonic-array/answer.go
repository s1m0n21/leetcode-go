/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/monotonic-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/28 10:07
 */

package _monotonic_array

func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	increasing := -1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}

		if increasing == -1 {
			if nums[i-1] < nums[i] {
				increasing = 1
			} else {
				increasing = 0
			}
		} else if (increasing == 1 && nums[i-1] > nums[i]) || (increasing == 0 && nums[i-1] < nums[i]) {
			return false
		}
	}

	return true
}
