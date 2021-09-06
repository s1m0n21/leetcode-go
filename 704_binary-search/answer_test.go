/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/6 9:37 上午
 */

package _binary_search

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{1, 2, 3, 4, 5}, 4}, 3},
		{input{[]int{1, 2, 3, 4, 5, 6}, 3}, 2},
		{input{[]int{1, 2, 3, 4, 5, 6}, 0}, -1},
		{input{[]int{}, 0}, -1},
	}

	for _, test := range tests {
		if actual := search(test.input.nums, test.input.target); actual != test.expect {
			t.Errorf("input = %+v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
