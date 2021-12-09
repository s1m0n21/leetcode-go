/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/single-number-iii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/9 9:27 AM
 */

package _single_number_iii

func singleNumber(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	lsb := xor & -xor
	type1, type2 := 0, 0
	for _, n := range nums {
		if n&lsb > 0 {
			type1 ^= n
		} else {
			type2 ^= n
		}
	}

	return []int{type1, type2}
}
