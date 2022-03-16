/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/reverse-bits/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/16 12:08 PM
 */

package _reverse_bits

func reverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; i < 32 && num > 0; i++ {
		rev |= num & 1 << (31 - i)
		num >>= 1
	}

	return rev
}
