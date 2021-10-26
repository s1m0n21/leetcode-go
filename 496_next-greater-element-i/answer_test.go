/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/26 9:44 上午
 */

package _next_greater_element_i

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		nums1, nums2 []int
	}

	tests := []struct {
		input  input
		expect []int
	}{
		{input{[]int{4, 1, 2}, []int{1, 3, 4, 2}}, []int{-1, 3, -1}},
		{input{[]int{2, 4}, []int{1, 2, 3, 4}}, []int{3, -1}},
	}

	for i, test := range tests {
		if actual := nextGreaterElement(test.input.nums1, test.input.nums2); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
