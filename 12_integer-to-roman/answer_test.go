/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/22 4:20 下午
 */

package _integer_to_roman

import "testing"

func TestAnswer(t *testing.T) {
	num := 1994

	t.Logf("answer = %s", intToRoman(num))
}
