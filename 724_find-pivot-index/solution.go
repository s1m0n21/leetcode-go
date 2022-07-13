/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/find-pivot-index/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/11 17:10
 */

package _find_pivot_index

func pivotIndex(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	curr := 0
	for i, n := range nums {
		if 2*curr+n == sum {
			return i
		}
		curr += n
	}

	return -1
}
