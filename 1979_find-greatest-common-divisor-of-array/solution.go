/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/find-greatest-common-divisor-of-array/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/16 10:17
 */

package _find_greatest_common_divisor_of_array

func findGCD(nums []int) int {
	var n1, n2 = nums[0], nums[0]
	for _, n := range nums {
		if n1 < n {
			n1 = n
		}
		if n2 > n {
			n2 = n
		}
	}

	for n1 != n2 {
		if n1 > n2 {
			n1 -= n2
		} else {
			n2 -= n1
		}
	}

	return n1
}
