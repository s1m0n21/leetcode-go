/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/20 2:11 下午
 */

package _longest_harmonious_subsequence

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{1, 1, 1, 1}, 0},
	}

	for i, test := range tests {
		if actual := findLHS(append([]int(nil), test.input...)); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
