/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/21 09:26
 */

package _n_repeated_element_in_size_2n_array

func repeatedNTimes(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		if m[n] {
			return n
		}
		m[n] = true
	}

	return -1
}
