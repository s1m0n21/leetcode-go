/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/7 3:28 PM
 */

package _triangle

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect int
	}{
		{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{[][]int{{-10}}, -10},
	}

	for i, test := range tests {
		if actual := minimumTotal(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
