/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/26 11:36 下午
 */

package _merge_sorted_array

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		nums1, nums2 []int
		n1, n2 int
	}

	tests := []struct{
		input input
		expect []int
	}{
		{input{[]int{1,2,3,0,0,0}, []int{2,5,6}, 3, 3}, []int{1,2,2,3,5,6}},
		{input{[]int{1}, nil, 1, 0}, []int{1}},
	}

	for _, test := range tests {
		actual := utils.CopySlice(test.input.nums1)
		if merge(actual, test.input.n1, test.input.nums2, test.input.n2); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("input = %+v, expect = %+v, acutal = %+v", test.input, test.expect, actual)
		}
	}
}
