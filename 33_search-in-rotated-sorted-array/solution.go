/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/search-in-rotated-sorted-array/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/9 11:46
 */

package _search_in_rotated_sorted_array

func search(nums []int, target int) int {
	var j, k int
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			k = len(nums) - i
			j = i
			break
		}
	}

	sorted := append(nums[j:], nums[0:j]...)
	l, r := 0, len(sorted)-1
	for l <= r {
		mid := (l + r) / 2
		if sorted[mid] < target {
			l = mid + 1
		} else if sorted[mid] > target {
			r = mid - 1
		} else {
			if sorted[mid] < sorted[k] {
				return mid + j
			}
			return mid - k
		}
	}

	return -1
}
