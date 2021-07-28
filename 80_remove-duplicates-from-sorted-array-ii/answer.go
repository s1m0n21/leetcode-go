/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/25 6:53 下午
 */

package _remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
