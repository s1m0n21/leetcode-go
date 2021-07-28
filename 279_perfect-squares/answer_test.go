/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/24 10:53 上午
 */

package _perfect_squares

import "testing"

func TestAnswer(t *testing.T) {
	n := 7168

	t.Logf("answer = %d", numSquares(n))
}
