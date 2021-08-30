/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/29 9:12 下午
 */

package interview_10_01_sorted_merge_lcci

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		A, B []int
		m, n int
	}

	tests := []struct {
		input  input
		expect []int
	}{
		{input{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3}, []int{1, 2, 2, 3, 5, 6}},
		{input{[]int{1, 2, 0, 0}, []int{1, 2}, 2, 2}, []int{1, 1, 2, 2}},
		{input{[]int{1}, []int{}, 1, 0}, []int{1}},
	}

	for _, test := range tests {
		c := utils.CopySlice(test.input.A)
		if merge(c, test.input.m, test.input.B, test.input.n); !reflect.DeepEqual(test.expect, c) {
			t.Errorf("input = %+v, expect = %#v, actual = %#v", test.input, test.expect, c)
		}
	}
}
