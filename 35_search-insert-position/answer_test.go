/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/15 9:27 PM
 */

package _search_insert_position

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{1, 3, 5, 6}, 5}, 2},
		{input{[]int{1, 3, 5, 6}, 2}, 1},
		{input{[]int{1, 3, 5, 6}, 7}, 4},
		{input{[]int{1, 3, 5, 6}, 0}, 0},
		{input{[]int{1}, 0}, 0},
	}

	for i, test := range tests {
		if actual := searchInsert(test.input.nums, test.input.target); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
