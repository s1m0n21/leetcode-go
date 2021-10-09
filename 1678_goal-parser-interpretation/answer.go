/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/goal-parser-interpretation/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/9 9:29 上午
 */

package _goal_parser_interpretation

func interpret(command string) string {
	var res []byte
	for i, s := range command {
		if s == 'G' {
			res = append(res, 'G')
		} else if s == '(' && command[i+1] == ')' {
			res = append(res, 'o')
		} else if s == '(' && command[i+1] == 'a' {
			res = append(res, 'a', 'l')
		}
	}

	return string(res)
}
