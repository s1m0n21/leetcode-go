/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/21 7:22 下午
 */

package _evaluate_reverse_polish_notation

import "testing"

func TestAnswer(t *testing.T) {
	tokens := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}

	t.Logf("answer = %d", evalRPN(tokens))
}
