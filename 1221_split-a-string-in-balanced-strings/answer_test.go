/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/7 1:36 下午
 */

package _split_a_string_in_balanced_strings

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"RLRRLLRLRL", 4},
		{"RLLLLRRRLR", 3},
		{"LLLLRRRR", 1},
		{"RLRRRLLRLL", 2},
	}

	for _, test := range tests {
		if actual := balancedStringSplit(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
