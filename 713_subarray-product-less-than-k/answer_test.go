/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/11 4:58 下午
 */

package _subarray_product_less_than_k

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
		{input{[]int{10, 5, 2, 6}, 100}, 8},
		{input{[]int{1, 2, 3}, 0}, 0},
	}

	for i, test := range tests {
		if actual := numSubarrayProductLessThanK(test.input.nums, test.input.k); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
