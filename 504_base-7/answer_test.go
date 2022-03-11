/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/11 5:46 PM
 */

package _base_7

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect string
	}{
		{100, "202"},
		{-7, "-10"},
	}

	for i, test := range tests {
		if actual := convertToBase7(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
