/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/6 11:14 上午
 */

package _ones_and_zeroes

import "testing"

func TestAnswer(t *testing.T) {
	strs := []string{
		"10",
		"0001",
		"111001",
		"1",
		"0",
	}
	m := 5
	n := 3

	t.Logf("answer = %d", findMaxForm(strs, m, n))
}
