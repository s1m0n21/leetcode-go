/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/number-of-1-bits/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/16 13:33
 */

package _number_of_1_bits

func hammingWeight(num uint32) int {
	var cnt int
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			cnt++
		}
	}
	return cnt
}
