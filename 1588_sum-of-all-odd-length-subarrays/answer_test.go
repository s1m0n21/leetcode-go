/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/29 12:33 上午
 */

package _sum_of_all_odd_length_subarrays

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 4, 2, 5, 3}, 58},
		{[]int{1, 2}, 3},
		{[]int{10, 11, 12}, 66},
	}

	for _, test := range tests {
		if actual := sumOddLengthSubarrays(test.input); actual != test.expect {
			t.Errorf("input = %#v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
