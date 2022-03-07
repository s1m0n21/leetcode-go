/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/minimum-remove-to-make-valid-parentheses/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/7 11:47 AM
 */

package _minimum_remove_to_make_valid_parentheses

func minRemoveToMakeValid(s string) string {
	var ans []rune
	var stack []int

	for _, b := range s {
		if b == '(' {
			stack = append(stack, len(ans))
		}
		if b == ')' {
			if len(stack) == 0 {
				continue
			}
			stack = stack[:len(stack)-1]
		}
		ans = append(ans, b)
	}

	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans[:stack[i]], ans[stack[i]+1:]...)
	}

	return string(ans)
}
