/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/single-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/13 10:40 AM
 */

package _single_number

func singleNumber(nums []int) int {
	ans := 0
	for _, n := range nums {
		ans ^= n
	}
	return ans
}
