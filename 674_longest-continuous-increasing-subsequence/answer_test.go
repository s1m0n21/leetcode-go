/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/14 3:14 下午
 */

package _longest_continuous_increasing_subsequence

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{2, 2, 2, 2, 2, 2}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 7}, 7},
	}

	for i, test := range tests {
		if actual := findLengthOfLCIS(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
