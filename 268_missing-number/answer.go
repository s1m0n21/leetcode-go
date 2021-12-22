/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/missing-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/6 11:37 下午
 */

package _missing_number

func missingNumber(nums []int) int {
	ans := len(nums)
	for i, n := range nums {
		ans ^= i ^ n
	}
	return ans
}
