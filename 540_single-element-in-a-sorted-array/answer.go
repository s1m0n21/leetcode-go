/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/10 8:21 PM
 */

package _single_element_in_a_sorted_array

func singleNonDuplicate(nums []int) int {
	var ans int
	for _, n := range nums {
		ans ^= n
	}
	return ans
}
