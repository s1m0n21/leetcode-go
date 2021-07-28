/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sort-colors/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/26 9:46 下午
 */

package _sort_colors

func sortColors(nums []int) {
	k, j := -1, len(nums)
	for i := 0; i < j; {
		if nums[i] == 0 {
			k++
			swap(nums, k, i)
		} else if nums[i] == 2 {
			j--
			swap(nums, i, j)
			continue
		}

		i++
	}
}

func swap(nums []int, a, b int) {
	n := nums[a]
	nums[a] = nums[b]
	nums[b] = n
}
