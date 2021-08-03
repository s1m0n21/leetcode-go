/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/4 12:54 上午
 */

package _shortest_unsorted_continuous_subarray

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{1}, 0},
	}

	for _, test := range tests {
		if actual := findUnsortedSubarray(test.input); actual != test.expect {
			t.Errorf("input = %#v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
