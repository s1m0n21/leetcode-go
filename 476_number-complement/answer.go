/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/number-complement/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/18 10:07 上午
 */

package _number_complement

import "strconv"

func findComplement(num int) int {
	var c string
	var bin string
	for ; num > 0; num /= 2 {
		lsb := num % 2
		bin = string('0' + byte(lsb)) + bin
	}

	for _, b := range bin {
		if b == '0' {
			c += "1"
		} else {
			c += "0"
		}
	}
	ans, _ := strconv.ParseInt(c, 2, 32)

	return int(ans)
}
