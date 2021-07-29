/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/7 1:32 上午
 */

package _two_sum

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}

	tests := []struct {
		input  input
		expect []int
	}{
		{input{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{input{[]int{2, 7, 11, 17}, 9}, []int{0, 1}},
		{input{[]int{3, 4, 2}, 5}, []int{0, 2}},
		{input{[]int{}, 1}, nil},
	}

	for _, test := range tests {
		if actual := twoSum(test.input.nums, test.input.target); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, actual)
		}
	}
}
