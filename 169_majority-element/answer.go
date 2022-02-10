/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/majority-element/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/10 10:05 AM
 */

package _majority_element

func majorityElement(nums []int) int {
	var major, cnt int

	for _, n := range nums {
		if cnt == 0 {
			major = n
		}

		if n == major {
			cnt++
		} else {
			cnt--
		}
	}

	return major
}
