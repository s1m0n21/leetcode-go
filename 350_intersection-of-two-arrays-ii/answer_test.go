/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/30 11:48 下午
 */

package _intersection_of_two_arrays_ii

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
		{input{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2, 2}},
		{input{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{4, 9}},
		{input{[]int{1, 2}, []int{1, 1}}, []int{1}},
		{input{[]int{0, 5, 8, 7, 2, 9, 7, 5}, []int{1, 4, 8, 9}}, []int{8, 9}},
	}

	for i, test := range tests {
		if actual := intersect(test.input.nums1, test.input.nums2); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
