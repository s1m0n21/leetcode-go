/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/add-digits/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/22 11:51
 */

package _add_digits

func addDigits(num int) int {
	return (num-1)%9 + 1
}
