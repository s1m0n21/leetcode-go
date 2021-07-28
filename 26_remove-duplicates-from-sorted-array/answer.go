/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/25 3:49 下午
 */

package _remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
