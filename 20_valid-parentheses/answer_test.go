/**
 * @Author         : s1m0n21
 * @Description    :
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/21 6:01 下午
 */

package _valid_parentheses

import "testing"

func TestAnswer(t *testing.T) {
	s := "{(})"

	t.Logf("answer = %t", isValid(s))
}
