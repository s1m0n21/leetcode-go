/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/3 10:03 AM
 */

package _sliding_window_maximum

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		nums []int
		k    int
	}

	tests := []struct {
		input  input
		expect []int
	}{
		{input{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []int{3, 3, 5, 5, 6, 7}},
		{input{[]int{1}, 1}, []int{1}},
		{input{[]int{-1, 1}, 1}, []int{-1, 1}},
		{input{[]int{9, 11}, 2}, []int{11}},
		{input{[]int{4, -2}, 2}, []int{4}},
	}

	for i, test := range tests {
		if actual := maxSlidingWindow(test.input.nums, test.input.k); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
