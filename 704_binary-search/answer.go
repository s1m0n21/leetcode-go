/**
 * @Author         : s1m0n21
 * @Description    : https://leetcode-cn.com/problems/binary-search/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/6 9:17 上午
 */

package _binary_search

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right - left) / 2 + left
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
