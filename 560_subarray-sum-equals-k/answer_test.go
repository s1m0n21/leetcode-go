/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/14 2:09 PM
 */

package _subarray_sum_equals_k

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
		{input{[]int{1, 1, 1}, 2}, 2},
		{input{[]int{1, 2, 3}, 3}, 2},
	}

	for i, test := range tests {
		if actual := subarraySum(test.input.nums, test.input.k); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
