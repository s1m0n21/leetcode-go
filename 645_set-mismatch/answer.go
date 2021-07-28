/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/set-mismatch/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/4 12:04 下午
 */

package _set_mismatch

func findErrorNums(nums []int) []int {
	var appeared = make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		appeared[nums[i]]++
	}

	missing, duplicate := 0, 0
	for i := 1; i <= len(nums); i++ {
		if appeared[i] == 0 {
			missing = i
		}
		if appeared[nums[i-1]] > 1 {
			duplicate = nums[i-1]
		}
		if missing > 0 && duplicate > 0 {
			return []int{duplicate, missing}
		}
	}

	return nil
}
