/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sort-colors/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/26 9:46 下午
 */

package _sort_colors

func sortColors(nums []int) {
	l, r := -1, len(nums)
	for i := 0; i < r; {
		if nums[i] == 0 {
			l++
			nums[l], nums[i] = nums[i], nums[l]
		} else if nums[i] == 2 {
			r--
			nums[r], nums[i] = nums[i], nums[r]
			continue
		}
		i++
	}
}
