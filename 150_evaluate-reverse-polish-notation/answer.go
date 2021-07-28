/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/21 7:03 下午
 */

package _evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			var after int
			switch tokens[i] {
			case "+":
				after = n1 + n2
			case "-":
				after = n1 - n2
			case "*":
				after = n1 * n2
			case "/":
				after = n1 / n2
			}

			stack = append(stack, after)
		default:
			n, _ := strconv.Atoi(tokens[i])
			stack = append(stack, n)
		}
	}

	return stack[0]
}
