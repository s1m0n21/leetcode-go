/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/4 10:46 上午
 */

package _reverse_integer

import "testing"

func TestAnswer(t *testing.T) {
	x := 1234567890
	t.Logf("x = %d, answer = %d", x, reverse(x))
}
