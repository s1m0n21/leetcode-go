/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/8 9:24 上午
 */

package _number_of_segments_in_a_string

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"Hello, my name is John", 5},
		{"Hello,     world", 2},
		{"     ", 0},
		{"     test", 1},
		{"   test1 test2    ", 2},
	}

	for i, test := range tests {
		if actual := countSegments(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
