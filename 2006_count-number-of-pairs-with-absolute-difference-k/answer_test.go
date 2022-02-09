/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/9 9:31 AM
 */

package _count_number_of_pairs_with_absolute_difference_k

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		nums []int
		k    int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{1, 2, 2, 1}, 1}, 4},
		{input{[]int{1, 3}, 0}, 0},
		{input{[]int{3, 2, 1, 5, 4}, 2}, 3},
	}

	for i, test := range tests {
		if actual := countKDifference(test.input.nums, test.input.k); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
