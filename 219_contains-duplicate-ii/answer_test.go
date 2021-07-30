/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/17 11:09 下午
 */

package _contains_duplicate_ii

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		nums []int
		k int
	}

	tests := []struct{
		input input
		expect bool
	}{
		{input{[]int{1,2,3,1}, 3}, true},
		{input{[]int{1,0,1,1}, 1}, true},
		{input{[]int{1,2,3,1,2,3}, 2}, false},
	}

	for _, test := range tests {
		if actual := containsNearbyDuplicate(test.input.nums, test.input.k); actual != test.expect {
			t.Errorf("input = %+v, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
