/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/roman-to-integer/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/4 10:52 上午
 */

package _roman_to_integer

var romanInt = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var out = 0
	for i := 0; i < len(s); i++ {
		if i < len(s) - 1 && romanInt[s[i]] < romanInt[s[i+1]] {
			out -= romanInt[s[i]]
			continue
		}
		out += romanInt[s[i]]
	}
	return out
}
