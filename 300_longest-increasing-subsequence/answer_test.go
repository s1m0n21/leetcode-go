/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/7 12:36 PM
 */

package _longest_increasing_subsequence

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}

	for i, test := range tests {
		if actual := lengthOfLIS(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, exepct = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
