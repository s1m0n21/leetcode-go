/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/1 9:10 上午
 */

package _compare_version_numbers

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		version1, version2 string
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{"1.0.0", "1.0"}, 0},
		{input{"1.0.1", "1.0"}, 1},
		{input{"1.0.1", "1.1"}, -1},
		{input{"1.0.00001", "1.0.1"}, 0},
	}

	for _, test := range tests {
		if actual := compareVersion(test.input.version1, test.input.version2); actual != test.expect {
			t.Errorf("input = %+v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
