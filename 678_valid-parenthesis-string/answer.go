/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/valid-parenthesis-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/15 11:44 下午
 */

package _valid_parenthesis_string

func checkValidString(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '*':
			stack = append(stack, s[i])
		case ')':
			if len(stack) > 0 {
				del := false
				for i := len(stack)-1; i >= 0; i-- {
					if stack[i] == '(' {
						stack = append(stack[:i], stack[i+1:]...)
						del = true
						break
					}
				}
				if !del {
					stack = stack[:len(stack)-1]
				}
			} else {
				return false
			}
		}
	}

	var c int
	for _, v := range stack {
		if v == '(' {
			c++
		} else {
			if c > 0 { c-- }
		}
	}

	return c == 0
}
