/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/valid-parentheses/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/21 5:52 下午
 */

package _valid_parentheses

var parentheses = map[byte]byte {
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	var stack = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if _, exist := parentheses[s[i]]; !exist {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || parentheses[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[0:len(stack)-1]
		}
	}

	return len(stack) == 0
}
