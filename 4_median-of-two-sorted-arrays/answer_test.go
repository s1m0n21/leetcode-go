/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/27 9:41 上午
 */

package _median_of_two_sorted_arrays

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		input  input
		expect float64
	}{
		{input{[]int{1, 3}, []int{2}}, 2},
		{input{[]int{1, 2}, []int{3, 4}}, 2.5},
		{input{[]int{0, 0}, []int{0, 0}}, 0},
		{input{[]int{}, []int{1}}, 1},
		{input{[]int{2}, []int{}}, 2},
	}

	for _, test := range tests {
		c := utils.CopySlice(test.input.nums1)
		if actual := findMedianSortedArrays(c, test.input.nums2); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("input = %+v, expect = %f, actual = %f", test.input, test.expect, actual)
		}
	}
}
