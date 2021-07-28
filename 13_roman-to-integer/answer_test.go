/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/4 10:57 上午
 */

package _roman_to_integer

import "testing"

func TestAnswer(t *testing.T) {
	s := "MCMXCIV"
	t.Logf("roman = %s, int = %d", s, romanToInt(s))
}
